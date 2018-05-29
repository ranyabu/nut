package nut

import (
	"container/list"
	"github.com/nut/i"
)

type LinkedMap struct {
	ks  *list.List
	kvs map[interface{}]interface{}
	nf  func(interface{}) bool
}

func NewLinkedMap() *LinkedMap {
	return &LinkedMap{
		ks:  list.New(),
		kvs: make(map[interface{}]interface{}),
	}
}

func NewLinkedMap2(nf func(interface{}) bool) *LinkedMap {
	return &LinkedMap{
		ks:  list.New(),
		kvs: make(map[interface{}]interface{}),
		nf:  nf,
	}
}

func (lm *LinkedMap) Len() int {
	if lm.ks != nil && lm.kvs != nil {
		return lm.ks.Len()
	} else {
		return 0
	}
}

func (lm *LinkedMap) ContainsKey(key interface{}) bool {
	if lm.isNil(key) {
		panic("key nil")
	}
	return lm.kvs[key] != nil
}

func (lm *LinkedMap) Get(key interface{}) interface{} {
	if lm.isNil(key) {
		panic("key nil")
	}
	return lm.kvs[key]
}

func (lm *LinkedMap) Put(key interface{}, value interface{}) interface{} {
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

func (lm *LinkedMap) PutIfAbsent(key interface{}, value interface{}) interface{} {
	if !lm.ContainsKey(key) {
		lm.Put(key, value)
		return nil
	}
	return lm.Get(key)
}

func (lm *LinkedMap) ComputeIfAbsent(key interface{}, siFunc func(key interface{}) interface{}) interface{} {
	if !lm.ContainsKey(key) {
		newV := siFunc(key)
		if !lm.isNil(newV) {
			return lm.Put(key, newV)
		}
	}
	return lm.Get(key)
}

func (lm *LinkedMap) ComputeIfPresent(key interface{}, biFunc func(key, value interface{}) interface{}) interface{} {
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

func (lm *LinkedMap) Remove(key interface{}) interface{} {
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

func (lm *LinkedMap) PutAll(m i.Map) {
	m.Foreach(func(kv ...interface{}) {
		lm.kvs[kv[0]] = kv[1]
	})
}

func (lm *LinkedMap) Clear() {
	lm.ks.Init()
	lm.kvs = make(map[interface{}]interface{})
}

func (lm *LinkedMap) Foreach(consumer func(...interface{})) {
	for e := lm.ks.Front(); e != nil; e = e.Next() {
		consumer(e.Value, lm.kvs[e.Value])
	}
}

func (si *LinkedMap) ForeachBreak(bk func(interface{}) bool, consumer func(...interface{})) interface{} {
	for key := range si.kvs {
		if b := bk(key); b {
			return key
		}
		consumer(key)
	}
	return nil
}

func (lm *LinkedMap) isNil(value interface{}) bool {
	return (lm.nf != nil && lm.nf(value)) || value == nil
}
