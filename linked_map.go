package nut

import (
	"container/list"
	"github.com/nut/iterf"
	"github.com/nut/util"
	"sync"
)

type linkedMap struct {
	ks   *list.List
	kvs  map[interface{}]interface{}
	lock sync.RWMutex
}

func NewLinkedMap() *linkedMap {
	return &linkedMap{
		ks:  list.New(),
		kvs: make(map[interface{}]interface{}),
	}
}

func (self *linkedMap) Len() int {
	self.lock.RLock()
	defer self.lock.RUnlock()
	
	return self.ks.Len()
}

func (self *linkedMap) ContainsKey(key interface{}) bool {
	self.lock.RLock()
	defer self.lock.RUnlock()
	
	return self.containsKey(key)
}

func (self *linkedMap) containsKey(key interface{}) bool {
	return util.IsNotNil(self.get(key))
}

func (self *linkedMap) Get(key interface{}) interface{} {
	self.lock.RLock()
	defer self.lock.RUnlock()
	
	return self.get(key)
}

func (self *linkedMap) get(key interface{}) interface{} {
	return self.kvs[key]
}

func (self *linkedMap) Put(key interface{}, value interface{}) interface{} {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	return self.put(key, value)
}

func (self *linkedMap) put(key interface{}, value interface{}) interface{} {
	if util.IsNil(self.kvs[key]) {
		self.ks.PushFront(key)
	} else {
		for e := self.ks.Front(); e != nil; e = e.Next() {
			self.ks.MoveToBack(e)
		}
	}
	self.kvs[key] = value
	return value
}

func (self *linkedMap) PutIfAbsent(key interface{}, value interface{}) interface{} {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	if !self.containsKey(key) {
		self.put(key, value)
		return nil
	}
	return self.get(key)
}

func (self *linkedMap) ComputeIfAbsent(key interface{}, siFunc func(key interface{}) interface{}) interface{} {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	if !self.containsKey(key) {
		newV := siFunc(key)
		if util.IsNotNil(newV) {
			return self.put(key, newV)
		}
	}
	return self.get(key)
}

func (self *linkedMap) ComputeIfPresent(key interface{}, biFunc func(key, value interface{}) interface{}) interface{} {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	if self.containsKey(key) {
		
		newV := biFunc(key, self.kvs[key])
		if util.IsNotNil(newV) {
			return self.put(key, newV)
		} else {
			self.remove(key)
			return nil
		}
	}
	return nil
}

func (self *linkedMap) Remove(key interface{}) interface{} {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	return self.remove(key)
}

func (self *linkedMap) remove(key interface{}) interface{} {
	for e := self.ks.Front(); e != nil; e = e.Next() {
		if key == e {
			self.ks.Remove(e)
			break
		}
	}
	value := self.kvs[key]
	delete(self.kvs, key)
	return value
}

func (self *linkedMap) PutAll(m iterf.Map) {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	m.Foreach(func(kv ...interface{}) {
		self.put(kv[0], kv[1])
	})
}

func (self *linkedMap) Clear() {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	self.ks.Init()
	self.kvs = make(map[interface{}]interface{})
}

func (self *linkedMap) Foreach(consumer func(...interface{})) {
	self.lock.RLock()
	defer self.lock.RUnlock()
	
	for e := self.ks.Front(); e != nil; e = e.Next() {
		consumer(e.Value, self.kvs[e.Value])
	}
}

func (self *linkedMap) ForeachBreak(bk func(...interface{}) bool, consumer func(...interface{})) interface{} {
	self.lock.RLock()
	defer self.lock.RUnlock()
	
	for e := self.ks.Front(); e != nil; e = e.Next() {
		if b := bk(e.Value); b {
			return e.Value
		}
		consumer(e.Value, self.kvs[e.Value])
	}
	
	return nil
}
