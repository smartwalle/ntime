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

func (b *bucket) Expiration() int64 {
	return atomic.LoadInt64(&b.expiration)
}

func (b *bucket) SetExpiration(expiration int64) bool {
	return atomic.SwapInt64(&b.expiration, expiration) != expiration
}

func (b *bucket) Add(t *timer) {
	if t == nil {
		return
	}
	b.mu.Lock()
	var ele = b.timers.PushBack(t)
	t.element = ele
	b.mu.Unlock()
}

func (b *bucket) remove(t *timer) {
	b.timers.Remove(t.element)
	t.element = nil
}

func (b *bucket) Remove(t *timer) {
	if t == nil {
		return
	}

	b.mu.Lock()
	b.remove(t)
	b.mu.Unlock()
}

func (b *bucket) Flush(fn func(t *timer)) {
	b.mu.Lock()

	for ele := b.timers.Front(); ele != nil; {
		var next = ele.Next()

		var t = ele.Value.(*timer)
		b.remove(t)

		fn(t)

		ele = next
	}
	b.SetExpiration(-1)

	b.mu.Unlock()
}
