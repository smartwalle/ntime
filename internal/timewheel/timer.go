package timewheel

import (
	"container/list"
	"sync/atomic"
)

type Timer interface {
	Stop()
}

type timer struct {
	expiration int64
	done       int32
	task       func()
	element    *list.Element
}

func newTimer(expiration int64, task func()) *timer {
	var t = &timer{}
	t.expiration = expiration
	t.task = task
	return t
}

func (t *timer) exec() {
	if atomic.CompareAndSwapInt32(&t.done, 0, 1) {
		t.task()
	}
}

func (t *timer) Stop() {
	if atomic.CompareAndSwapInt32(&t.done, 0, 1) {
	}
}
