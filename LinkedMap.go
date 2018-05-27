package nut

import (
	"container/list"
	"github.com/nut/i"
)

type LinkedMap struct {
	ks  *list.List
	kvs map[interface{}]interface{}
}

func NewLinkedMap() *LinkedMap {
	return &LinkedMap{
		ks:  list.New(),
		kvs: make(map[interface{}]interface{}),
	}
}

func (lm *LinkedMap) Len() int {
	return lm.ks.Len()
}

func (lm *LinkedMap) ContainsKey(key interface{}) bool {
	return lm.kvs[key] != nil
}

func (lm *LinkedMap) ContainsValue(value interface{}) bool {
	for _, v := range lm.kvs {
		if v == value {
			return true
		}
	}
	return false
}

func (lm *LinkedMap) Get(key interface{}) interface{} {
	return lm.kvs[key]
}

func (lm *LinkedMap) Put(key interface{}, value interface{}) interface{} {
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
	if !lm.ContainsKey(key) && value != nil {
		lm.Put(key, value)
		return nil
	}
	return value
}

func (lm *LinkedMap) ComputeIfAbsent(key interface{}, siFunc func(key interface{}) interface{}) interface{} {
	if !lm.ContainsKey(key) {
		newV := siFunc(key)
		if newV != nil {
			return lm.Put(key, newV)
		}
	}
	return nil
}

func (lm *LinkedMap) ComputeIfPresent(key interface{}, biFunc func(key, value interface{}) interface{}) interface{} {
	if value := lm.Get(key); value != nil {
		newV := biFunc(key, lm.kvs[key])
		if newV != nil {
			return lm.Put(key, newV)
		} else {
			lm.Remove(key)
			return nil
		}
	}
	return nil
}

func (lm *LinkedMap) Remove(key interface{}) interface{} {
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

func (lm *LinkedMap) Foreach(consumer func(...interface{})) {
	for e := lm.ks.Front(); e != nil; e = e.Next() {
		consumer(e.Value, lm.kvs[e.Value])
	}
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
