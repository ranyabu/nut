package nut

import (
	"container/list"
	"github.com/nut/iterf"
	"github.com/nut/common"
)

type linkedMap struct {
	ks  *list.List
	kvs map[interface{}]interface{}
}

func NewLinkedMap() *linkedMap {
	return &linkedMap{
		ks:  list.New(),
		kvs: make(map[interface{}]interface{}),
	}
}

func (lm *linkedMap) Len() int {
	return lm.ks.Len()
}

func (lm *linkedMap) ContainsKey(key interface{}) bool {
	
	return common.IsNotNil(lm.kvs[key])
}

func (lm *linkedMap) Get(key interface{}) interface{} {
	return lm.kvs[key]
}

func (lm *linkedMap) Put(key interface{}, value interface{}) interface{} {
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
		if common.IsNotNil(newV) {
			return lm.Put(key, newV)
		}
	}
	return lm.Get(key)
}

func (lm *linkedMap) ComputeIfPresent(key interface{}, biFunc func(key, value interface{}) interface{}) interface{} {
	if lm.ContainsKey(key) {
		newV := biFunc(key, lm.kvs[key])
		if common.IsNotNil(newV) {
			return lm.Put(key, newV)
		} else {
			lm.Remove(key)
			return nil
		}
	}
	return nil
}

func (lm *linkedMap) Remove(key interface{}) interface{} {
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

func (lm *linkedMap) ForeachBreak(bk func(...interface{}) bool, consumer func(...interface{})) interface{} {
	
	for e := lm.ks.Front(); e != nil; e = e.Next() {
		if b := bk(e.Value); b {
			return e.Value
		}
		consumer(e.Value, lm.kvs[e.Value])
	}
	
	return nil
}
