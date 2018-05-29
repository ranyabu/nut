package nut

import (
	"container/list"
	"github.com/nut/iterf"
)

type linkedMap struct {
	ks   *list.List
	kvs  map[interface{}]interface{}
	nf   func(interface{}) bool
}

func NewLinkedMap() *linkedMap {
	return &linkedMap{
		ks:   list.New(),
		kvs:  make(map[interface{}]interface{}),
	}
}

func NewLinkedMap2(nf func(interface{}) bool) *linkedMap {
	return &linkedMap{
		ks:  list.New(),
		kvs: make(map[interface{}]interface{}),
		nf:  nf,
	}
}

func (lm *linkedMap) Len() int {
	return lm.ks.Len()
}

func (lm *linkedMap) ContainsKey(key interface{}) bool {
	if lm.isNil(key) {
		panic("key nil")
	}
	return lm.kvs[key] != nil
}

func (lm *linkedMap) Get(key interface{}) interface{} {
	if lm.isNil(key) {
		panic("key nil")
	}
	return lm.kvs[key]
}

func (lm *linkedMap) Put(key interface{}, value interface{}) interface{} {
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

func (lm *linkedMap) PutIfAbsent(key interface{}, value interface{}) interface{} {
	if !lm.ContainsKey(key) {
		lm.Put(key, value)
		return nil
	}
	return lm.Get(key)
}

func (lm *linkedMap) ComputeIfAbsent(key interface{}, siFunc func(key interface{}) interface{}) interface{} {
	if !lm.ContainsKey(key) {
		newV := siFunc(key)
		if !lm.isNil(newV) {
			return lm.Put(key, newV)
		}
	}
	return lm.Get(key)
}

func (lm *linkedMap) ComputeIfPresent(key interface{}, biFunc func(key, value interface{}) interface{}) interface{} {
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

func (lm *linkedMap) Remove(key interface{}) interface{} {
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

func (lm *linkedMap) PutAll(m iterf.Map) {
	m.Foreach(func(kv ...interface{}) {
		lm.kvs[kv[0]] = kv[1]
	})
}

func (lm *linkedMap) Clear() {
	lm.ks.Init()
	lm.kvs = make(map[interface{}]interface{})
}

func (lm *linkedMap) Foreach(consumer func(...interface{})) {
	for e := lm.ks.Front(); e != nil; e = e.Next() {
		consumer(e.Value, lm.kvs[e.Value])
	}
}

func (si *linkedMap) ForeachBreak(bk func(...interface{}) bool, consumer func(...interface{})) interface{} {
	for key := range si.kvs {
		if b := bk(key); b {
			return key
		}
		consumer(key)
	}
	return nil
}

func (lm *linkedMap) isNil(value interface{}) bool {
	return (lm.nf != nil && lm.nf(value)) || value == nil
}
