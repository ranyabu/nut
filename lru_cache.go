package nut

import (
	"container/list"
	"github.com/nut/util"
	"sync"
)

type lruCache struct {
	ks   *list.List
	kvs  map[interface{}]interface{}
	max  int
	lock sync.RWMutex
}

func NewLRUCache(max int) *lruCache {
	return &lruCache{
		ks:  list.New(),
		kvs: make(map[interface{}]interface{}),
		max: max,
	}
}

func (self *lruCache) Len() int {
	self.lock.RLock()
	defer self.lock.RUnlock()
	
	return self.ks.Len()
}

func (self *lruCache) Contains(key interface{}) bool {
	self.lock.RLock()
	defer self.lock.RUnlock()
	
	return self.contains(key)
}

func (self *lruCache) contains(key interface{}) bool {
	return util.IsNotNil(self.kvs[key])
}

func (self *lruCache) Add(key, value interface{}) {
	self.AddIfAbsent(key, value)
}

func (self *lruCache) AddIfAbsent(key, value interface{}) interface{} {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	if self.contains(key) {
		for e := self.ks.Front(); e != nil; e = e.Next() {
			self.ks.MoveToBack(e)
			break
		}
		return self.kvs[key]
	} else {
		if self.ks.Len() == self.max {
			delete(self.kvs, self.ks.Front().Value)
			self.ks.Remove(self.ks.Front())
		}
		
		self.ks.PushBack(key)
		self.kvs[key] = value
		return nil
	}
}

func (self *lruCache) Remove(key interface{}) bool {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	return self.remove(key)
}

func (self *lruCache) remove(key interface{}) bool {
	if self.contains(key) {
		for e := self.ks.Front(); e != nil; e = e.Next() {
			self.ks.Remove(e)
			break
		}
		delete(self.kvs, key)
		return true
	}
	
	return false
}

func (self *lruCache) Foreach(consumer func(...interface{})) {
	self.lock.RLock()
	defer self.lock.RUnlock()
	
	for e := self.ks.Front(); e != nil; e = e.Next() {
		consumer(e.Value, self.kvs[e.Value])
	}
}

func (self *lruCache) ForeachBreak(bk func(...interface{}) bool, consumer func(...interface{})) interface{} {
	self.lock.RLock()
	defer self.lock.RUnlock()
	
	for key := range self.kvs {
		if b := bk(key); b {
			return key
		}
		consumer(key, self.kvs[key])
	}
	return nil
}

func (self *lruCache) Clear() {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	self.ks.Init()
	self.kvs = make(map[interface{}]interface{})
}
