package timewheel

import (
	"errors"
	"github.com/smartwalle/queue/delay"
	"sync"
	"time"
)

type TimeWheel struct {
	currentTime int64 // milliseconds
	tick        int64 // milliseconds
	size        int64
	buckets     []*Bucket
	queue       delay.Queue[*Bucket]
	runOnce     sync.Once
}

func New(tick time.Duration, size int64) *TimeWheel {
	var mTick = int64(tick / time.Millisecond)
	if mTick <= 0 {
		panic(errors.New("tick must be greater than or equal to 1ms"))
	}

	if size <= 0 {
		panic(errors.New("size must be greater than or equal to 1"))
	}

	var buckets = make([]*Bucket, size)

	for i := range buckets {
		buckets[i] = newBucket()
	}

	var tw = &TimeWheel{}
	tw.currentTime = time.Now().UnixMilli()
	tw.tick = mTick
	tw.size = size
	tw.buckets = buckets
	tw.queue = delay.New[*Bucket](
		delay.WithTimeUnit(time.Millisecond),
		delay.WithTimeProvider(func() int64 {
			return time.Now().UnixMilli()
		}),
	)

	return tw
}

func (this *TimeWheel) Run() {
	this.runOnce.Do(func() {
		go this.run()
	})
}

func (this *TimeWheel) run() {
	for {
		var bucket, expiration = this.queue.Dequeue()
		bucket.Flush(expiration)
	}
}

func (this *TimeWheel) After(delay time.Duration, fn func()) *Timer {
	var expiration = time.Now().Add(delay).UnixMilli()
	var timer = newTimer(expiration, fn)

	var index = expiration % this.size
	var bucket = this.buckets[index]

	bucket.Add(timer)

	this.queue.Enqueue(bucket, expiration)

	return timer
}
