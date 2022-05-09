package timewheel

import (
	"container/list"
	"sync"
)

type Bucket struct {
	mu sync.Mutex
	//timers priority.Queue[*Timer]
	timers *list.List
}

func newBucket() *Bucket {
	var b = &Bucket{}
	//b.timers = priority.New[*Timer]()
	b.timers = list.New()
	return b
}

func (this *Bucket) Add(t *Timer) {
	if t == nil {
		return
	}
	this.mu.Lock()
	//var ele = this.timers.Enqueue(t, t.expiration)
	var ele = this.timers.PushBack(t)
	t.element = ele
	this.mu.Unlock()
}

func (this *Bucket) remove(t *Timer) {
	this.timers.Remove(t.element)
	t.element = nil
}

func (this *Bucket) Remove(t *Timer) {
	if t == nil {
		return
	}

	this.mu.Lock()
	this.remove(t)
	this.mu.Unlock()
}

func (this *Bucket) Flush(now int64) {
	this.mu.Lock()

	//for timer, _, _, ok := this.timers.Peek(now); ok; {
	//	timer.exec()
	//
	//	timer, _, _, ok = this.timers.Peek(now)
	//}

	for ele := this.timers.Front(); ele != nil; {
		var next = ele.Next()

		var timer = ele.Value.(*Timer)
		this.remove(timer)

		timer.exec()

		ele = next
	}

	this.mu.Unlock()
}
