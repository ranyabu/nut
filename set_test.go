package nut

import (
	"testing"
	"sync"
)

func TestSet(t *testing.T) {
	var lock sync.RWMutex
	lock.RLock()
	lock.Lock()
	//lock.RUnlock()
	//lock.Unlock()
}

