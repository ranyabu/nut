package nut

import (
	"container/list"
)

type LRUMap struct {
	ks  *list.List
	kvs map[interface{}]interface{}
	nf  func(interface{}) bool
	max int
}

var defaultValue = []byte{0}

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

func (lm *LRUMap) Contains(value interface{}) bool {
	if lm.isNil(value) {
		panic("value nil")
	}
	
	return lm.isNil(lm.kvs[value])
	
}

func (lm *LRUMap) Add(value interface{}) {
	lm.AddIfAbsent(value)
}

func (lm *LRUMap) AddIfAbsent(value interface{}) interface{} {
	if lm.isNil(value) {
		panic("value nil")
	}
	
	if lm.Contains(value) {
		for e := lm.ks.Front(); e != nil; e = e.Next() {
			lm.ks.MoveToBack(e)
			break
		}
		return value
	} else {
		if lm.ks.Len() == lm.max {
			delete(lm.kvs, lm.ks.Front().Value)
			lm.ks.Remove(lm.ks.Front())
		}
		
		lm.ks.PushBack(value)
		lm.kvs[value] = defaultValue
		return nil
	}
}

func (lm *LRUMap) Remove(value interface{}) bool {
	if lm.isNil(value) {
		panic("value nil")
	}
	
	if lm.Contains(value) {
		for e := lm.ks.Front(); e != nil; e = e.Next() {
			lm.ks.Remove(e)
			break
		}
		delete(lm.kvs, value)
		return true
	}
	
	return false
}

func (lm *LRUMap) Foreach(consumer func(...interface{})) {
	for e := lm.ks.Front(); e != nil; e = e.Next() {
		consumer(e.Value)
	}
}

func (lm *LRUMap) ForeachBreak(bk func(interface{}) bool, consumer func(...interface{})) interface{} {
	for value := range lm.kvs {
		if b := bk(value); b {
			return value
		}
		consumer(value)
	}
	return nil
}

func (lm *LRUMap) Clear() {
	lm.ks.Init()
	lm.kvs = make(map[interface{}]interface{})
}

func (lm *LRUMap) isNil(value interface{}) bool {
	return (lm.nf != nil && lm.nf(value)) || value == nil
}
