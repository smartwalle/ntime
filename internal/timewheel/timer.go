package timewheel

import (
	"container/list"
	"sync/atomic"
)

type Timer struct {
	expiration int64
	done       int32
	task       func()
	//element    priority.Element
	element *list.Element
}

func newTimer(expiration int64, task func()) *Timer {
	var t = &Timer{}
	t.expiration = expiration
	t.task = task
	return t
}

func (this *Timer) exec() {
	if atomic.CompareAndSwapInt32(&this.done, 0, 1) {
		go this.task()
	}
}

func (this *Timer) Stop() {
	if atomic.CompareAndSwapInt32(&this.done, 0, 1) {
	}
}
