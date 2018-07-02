package locks

import (
	"time"
	"sync"
)

type keyLock struct {
	smt  sync.Mutex
	lock map[interface{}]struct{}
}

func (lk *keyLock) Lock(key interface{}) bool {
	lk.smt.Lock()
	defer lk.smt.Unlock()
	if lk.lock[key] == struct{}{} {
		lk.lock[key] = struct{}{}
		return true
	} else {
		return false
	}
}

func (lk *keyLock) Unlock(key interface{}) {
	delete(lk.lock, key)
}

func (lk *keyLock) TryLock(key interface{}, timeout time.Duration) bool {
	now := time.Now().Nanosecond()
	flag := true
	for ; !lk.Lock(key); {
		if int64(time.Now().Nanosecond()-now) >= timeout.Nanoseconds() {
			flag = false
			break
		}
		time.Sleep(time.Millisecond)
	}
	return flag
}

func (lk *keyLock) isLock(key interface{}) bool {
	return lk.lock[key] == struct{}{}
}
