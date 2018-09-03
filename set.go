package nut

import (
	"github.com/nut/iterf"
	"github.com/nut/util"
	"sync"
)

type set struct {
	kvs map[interface{}]struct{}
	lock sync.RWMutex
}

func NewSet() *set {
	return &set{kvs: make(map[interface{}]struct{})}
}

func (self *set) Len() int {
	self.lock.RLock()
	defer self.lock.RUnlock()
	
	return len(self.kvs)
}

func (self *set) Contains(value interface{}) bool {
	self.lock.RLock()
	defer self.lock.RUnlock()
	
	return self.contains(value)
}

func (self *set) contains(value interface{}) bool {
	return util.IsNotNil(self.kvs[value])
}

func (self *set) Add(value interface{}) bool {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	return self.add(value)
}

func (self *set) add(value interface{}) bool {
	if self.contains(value) {
		return false
	}
	self.kvs[value] = struct{}{}
	return true
}

func (self *set) Remove(value interface{}) bool {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	return self.remove(value)
}

func (self *set) remove(value interface{}) bool {
	if !self.contains(value) {
		return false
	}
	delete(self.kvs, value)
	return true
}

func (self *set) ContainsAll(set iterf.Set) bool {
	self.lock.RLock()
	defer self.lock.RUnlock()
	
	set.ForeachBreak(func(value ...interface{}) bool {
		return util.IsNil(self.kvs[value[0]])
	}, func(consumer ...interface{}) {
		// pass
	})
	return true
}

func (self *set) AddAll(set iterf.Set) bool {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	set.Foreach(func(value ...interface{}) {
		self.add(value[0])
	})
	return true
}

func (self *set) RetainAll(set iterf.Set) bool {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	modified := false
	set.Foreach(func(value ...interface{}) {
		if !self.contains(value[0]) {
			self.remove(value[0])
			modified = true
		}
	})
	return modified
}

func (self *set) RemoveAll(set iterf.Set) bool {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	set.Foreach(func(value ...interface{}) {
		self.remove(value[0])
	})
	return true
}

func (self *set) Clear() {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	self.kvs = make(map[interface{}]struct{})
}

func (self *set) Foreach(consumer func(...interface{})) {
	self.lock.RLock()
	defer self.lock.RUnlock()
	
	for key := range self.kvs {
		consumer(key)
	}
}

func (self *set) ForeachBreak(bk func(interface{}) bool, consumer func(...interface{})) interface{} {
	self.lock.RLock()
	defer self.lock.RUnlock()
	
	for key := range self.kvs {
		if b := bk(key); b {
			return key
		}
		consumer(key)
	}
	return nil
}
