package timewheel

import (
	"container/list"
	"sync"
	"sync/atomic"
)

type bucket struct {
	expiration int64

	mu     sync.Mutex
	timers *list.List
}

func newBucket() *bucket {
	var b = &bucket{}
	b.timers = list.New()
	return b
}

func (this *bucket) Expiration() int64 {
	return atomic.LoadInt64(&this.expiration)
}

func (this *bucket) SetExpiration(expiration int64) bool {
	return atomic.SwapInt64(&this.expiration, expiration) != expiration
}

func (this *bucket) Add(t *timer) {
	if t == nil {
		return
	}
	this.mu.Lock()
	var ele = this.timers.PushBack(t)
	t.element = ele
	this.mu.Unlock()
}

func (this *bucket) remove(t *timer) {
	this.timers.Remove(t.element)
	t.element = nil
}

func (this *bucket) Remove(t *timer) {
	if t == nil {
		return
	}

	this.mu.Lock()
	this.remove(t)
	this.mu.Unlock()
}

func (this *bucket) Flush(f func(t *timer)) {
	this.mu.Lock()

	for ele := this.timers.Front(); ele != nil; {
		var next = ele.Next()

		var t = ele.Value.(*timer)
		this.remove(t)

		f(t)

		ele = next
	}

	this.mu.Unlock()

	this.SetExpiration(-1)
}
