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

	if size <= 1 {
		panic(errors.New("size must be greater than or equal to 2"))
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

func (w *timeWheel) Run() {
	w.runOnce.Do(func() {
		go w.run()
	})
}

func (w *timeWheel) Close() {
	w.queue.Close()
}

func (w *timeWheel) run() {
	for {
		var b, expiration = w.queue.Dequeue()
		if expiration < 0 {
			return
		}
		w.clock(expiration)
		b.Flush(w.addOrRun)
	}
}

func (w *timeWheel) clock(expiration int64) {
	var nTime = atomic.LoadInt64(&w.currentTime)
	if expiration >= nTime+w.tick {
		nTime = truncate(expiration, w.tick)
		atomic.StoreInt64(&w.currentTime, nTime)

		var overflow = atomic.LoadPointer(&w.overflow)
		if overflow != nil {
			(*timeWheel)(overflow).clock(nTime)
		}
	}
}

func (w *timeWheel) add(t *timer) bool {
	var nTime = atomic.LoadInt64(&w.currentTime)

	if t.expiration < nTime+w.tick {
		return false
	} else if t.expiration < nTime+w.interval {
		var group = t.expiration / w.tick
		var index = group % w.size

		var b = w.buckets[index]
		b.Add(t)

		if b.SetExpiration(group * w.tick) {
			w.queue.Enqueue(b, b.Expiration())
		}

		return true
	} else {
		var overflow = atomic.LoadPointer(&w.overflow)
		if overflow == nil {
			atomic.CompareAndSwapPointer(
				&w.overflow,
				nil,
				unsafe.Pointer(
					newTimeWheel(
						w.interval,
						w.size,
						nTime,
						w.queue,
					),
				),
			)
			overflow = atomic.LoadPointer(&w.overflow)
		}

		return (*timeWheel)(overflow).add(t)
	}
}

func (w *timeWheel) addOrRun(t *timer) {
	if !w.add(t) {
		go t.exec()
	}
}

func (w *timeWheel) AfterFunc(delay time.Duration, fn func()) Timer {
	var expiration = time.Now().Add(delay).UnixMilli()
	var t = newTimer(expiration, fn)

	w.addOrRun(t)

	return t
}

func truncate(x, m int64) int64 {
	if m <= 0 {
		return x
	}
	return x - x%m
}
