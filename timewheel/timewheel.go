package timewheel

import (
	"errors"
	"github.com/smartwalle/queue/delay"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

type TimeWheel interface {
	Run()

	Close()

	AfterFunc(delay time.Duration, fn func()) Timer
}

type timeWheel struct {
	currentTime int64 // milliseconds
	tick        int64 // milliseconds
	interval    int64 // milliseconds
	size        int64
	buckets     []*bucket
	queue       delay.Queue[*bucket]
	runOnce     sync.Once
	overflow    unsafe.Pointer
}

func New(tick time.Duration, size int64) TimeWheel {
	var mTick = int64(tick / time.Millisecond)
	if mTick <= 0 {
		panic(errors.New("tick must be greater than or equal to 1ms"))
	}

	if size <= 0 {
		panic(errors.New("size must be greater than or equal to 1"))
	}

	var queue = delay.New[*bucket](
		delay.WithTimeUnit(time.Millisecond),
		delay.WithTimeProvider(func() int64 {
			return time.Now().UnixMilli()
		}),
	)

	return newTimeWheel(mTick, size, time.Now().UnixMilli(), queue)
}

func newTimeWheel(tick int64, size int64, startTime int64, queue delay.Queue[*bucket]) *timeWheel {
	var buckets = make([]*bucket, size)
	for i := range buckets {
		buckets[i] = newBucket()
	}

	var tw = &timeWheel{}
	tw.currentTime = truncate(startTime, tick)
	tw.tick = tick
	tw.interval = tick * size
	tw.size = size
	tw.buckets = buckets
	tw.queue = queue
	return tw
}

func (this *timeWheel) Run() {
	this.runOnce.Do(func() {
		go this.run()
	})
}

func (this *timeWheel) Close() {
	this.queue.Close()
}

func (this *timeWheel) run() {
	for {
		var b, expiration = this.queue.Dequeue()
		if expiration < 0 {
			return
		}
		this.clock(expiration)
		b.Flush(this.addOrRun)
	}
}

func (this *timeWheel) clock(expiration int64) {
	var nTime = atomic.LoadInt64(&this.currentTime)
	if expiration >= nTime+this.tick {
		nTime = truncate(expiration, this.tick)
		atomic.StoreInt64(&this.currentTime, nTime)

		var overflow = atomic.LoadPointer(&this.overflow)
		if overflow != nil {
			(*timeWheel)(overflow).clock(nTime)
		}
	}
}

func (this *timeWheel) add(t *timer) bool {
	var nTime = atomic.LoadInt64(&this.currentTime)

	if t.expiration < nTime+this.tick {
		return false
	} else if t.expiration < nTime+this.interval {
		var group = t.expiration / this.tick
		var index = group % this.size

		var b = this.buckets[index]
		b.Add(t)

		if b.SetExpiration(group * this.tick) {
			this.queue.Enqueue(b, b.Expiration())
		}

		return true
	} else {
		var overflow = atomic.LoadPointer(&this.overflow)
		if overflow == nil {
			atomic.CompareAndSwapPointer(
				&this.overflow,
				nil,
				unsafe.Pointer(newTimeWheel(this.interval, this.size, nTime, this.queue)),
			)
			overflow = atomic.LoadPointer(&this.overflow)
		}

		return (*timeWheel)(overflow).add(t)
	}
}

func (this *timeWheel) addOrRun(t *timer) {
	if !this.add(t) {
		go t.exec()
	}
}

func (this *timeWheel) AfterFunc(delay time.Duration, fn func()) Timer {
	var expiration = time.Now().Add(delay).UnixMilli()
	var t = newTimer(expiration, fn)

	this.addOrRun(t)

	return t
}

func truncate(x, m int64) int64 {
	if m <= 0 {
		return x
	}
	return x - x%m
}
