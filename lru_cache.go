package nut

import (
	"container/list"
	"github.com/nut/i"
)

type LRUMap struct {
	ks  *list.List
	kvs map[interface{}]interface{}
	nf  func(interface{}) bool
	max int
}

func NewLRUMap(max int) *LRUMap {
	return &LRUMap{
		ks:  list.New(),
		kvs: make(map[interface{}]interface{}),
		max: max,
	}
}

func NewLRUMap2(max int, nf func(interface{}) bool) *LRUMap {
	return &LRUMap{
		ks:  list.New(),
		kvs: make(map[interface{}]interface{}),
		nf:  nf,
		max: max,
	}
}

func (lm *LRUMap) Len() int {
	if lm.ks != nil && lm.kvs != nil {
		return lm.ks.Len()
	} else {
		return 0
	}
}

func (lm *LRUMap) Get(key interface{}) interface{} {
	if lm.isNil(key) {
		panic("key nil")
	}
	return lm.kvs[key]
}

func (lm *LRUMap) Put(key interface{}, value interface{}) interface{} {
	if lm.isNil(key) || lm.isNil(value) {
		panic("key or value nil")
	}
	
	if lm.ContainsKey(key) {
		for e := lm.ks.Front(); e != nil; e = e.Next() {
			lm.ks.MoveToBack(e)
		}
	} else {
		lm.ks.PushBack(key)
	}
	
	lm.kvs[key] = value
	return value
}

func (lm *LRUMap) PutIfAbsent(key interface{}, value interface{}) interface{} {
	if !lm.ContainsKey(key) {
		lm.Put(key, value)
		return nil
	}
	return lm.Get(key)
}

func (lm *LRUMap) ComputeIfAbsent(key interface{}, siFunc func(key interface{}) interface{}) interface{} {
	if !lm.ContainsKey(key) {
		newV := siFunc(key)
		if !lm.isNil(newV) {
			return lm.Put(key, newV)
		}
	}
	return lm.Get(key)
}

func (lm *LRUMap) ComputeIfPresent(key interface{}, biFunc func(key, value interface{}) interface{}) interface{} {
	if lm.ContainsKey(key) {
		newV := biFunc(key, lm.kvs[key])
		if !lm.isNil(newV) {
			return lm.Put(key, newV)
		} else {
			lm.Remove(key)
			return nil
		}
	}
	return nil
}

func (lm *LRUMap) Remove(key interface{}) interface{} {
	if lm.isNil(key) {
		panic("key nil")
	}
	
	for e := lm.ks.Front(); e != nil; e = e.Next() {
		if key == e {
			lm.ks.Remove(e)
			break
		}
	}
	value := lm.kvs[key]
	delete(lm.kvs, key)
	return value
}

func (lm *LRUMap) Foreach(consumer func(...interface{})) {
	for e := lm.ks.Front(); e != nil; e = e.Next() {
		consumer(e.Value, lm.kvs[e.Value])
	}
}

func (lm *LRUMap) PutAll(m i.Map) {
	m.Foreach(func(kv ...interface{}) {
		lm.kvs[kv[0]] = kv[1]
	})
}

func (lm *LRUMap) Clear() {
	lm.ks.Init()
	lm.kvs = make(map[interface{}]interface{})
}

func (lm *LRUMap) isNil(value interface{}) bool {
	return (lm.nf != nil && lm.nf(value)) || value == nil
}
