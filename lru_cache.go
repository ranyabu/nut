package nut

import (
	"container/list"
	"github.com/nut/util"
	"sync"
)

type lruCache struct {
	ks   *list.List
	kvs  map[interface{}]struct{}
	max  int
	lock sync.RWMutex
}

func NewLRUCache(max int) *lruCache {
	return &lruCache{
		ks:  list.New(),
		kvs: make(map[interface{}]struct{}),
		max: max,
	}
}

func (self *lruCache) Len() int {
	self.lock.RLock()
	defer self.lock.RUnlock()
	
	return self.ks.Len()
}

func (self *lruCache) Contains(value interface{}) bool {
	self.lock.RLock()
	defer self.lock.RUnlock()
	
	return self.contains(value)
}

func (self *lruCache) contains(value interface{}) bool {
	return util.IsNotNil(self.kvs[value])
}

func (self *lruCache) Add(value interface{}) {
	self.AddIfAbsent(value)
}

func (self *lruCache) AddIfAbsent(value interface{}) interface{} {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	if self.contains(value) {
		for e := self.ks.Front(); e != nil; e = e.Next() {
			self.ks.MoveToBack(e)
			break
		}
		return value
	} else {
		if self.ks.Len() == self.max {
			delete(self.kvs, self.ks.Front().Value)
			self.ks.Remove(self.ks.Front())
		}
		
		self.ks.PushBack(value)
		self.kvs[value] = struct{}{}
		return nil
	}
}

func (self *lruCache) Remove(value interface{}) bool {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	return self.remove(value)
}

func (self *lruCache) remove(value interface{}) bool {
	if self.contains(value) {
		for e := self.ks.Front(); e != nil; e = e.Next() {
			self.ks.Remove(e)
			break
		}
		delete(self.kvs, value)
		return true
	}
	
	return false
}

func (self *lruCache) Foreach(consumer func(...interface{})) {
	self.lock.RLock()
	defer self.lock.RUnlock()
	
	for e := self.ks.Front(); e != nil; e = e.Next() {
		consumer(e.Value)
	}
}

func (self *lruCache) ForeachBreak(bk func(...interface{}) bool, consumer func(...interface{})) interface{} {
	self.lock.RLock()
	defer self.lock.RUnlock()
	
	for value := range self.kvs {
		if b := bk(value); b {
			return value
		}
		consumer(value)
	}
	return nil
}

func (self *lruCache) Clear() {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	self.ks.Init()
	self.kvs = make(map[interface{}]struct{})
}
