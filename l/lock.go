package l

import (
	"time"
)

type lock struct {
	ch chan struct{}
}

func NewLock() *lock {
	return &lock{ch: make(chan struct{}, 1)}
}

func (lk *lock) Lock() {
	<-lk.ch
}

func (lk *lock) Unlock() {
	select {
	case lk.ch <- struct{}{}:
	default:
		panic("unlock")
	}
}
func (lk *lock) TryLock(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case <-lk.ch:
		timer.Stop()
		return true
	case <-time.After(timeout):
	}
	return false
}

func (lk *lock) isLock() bool {
	return len(lk.ch) > 0
	
}
