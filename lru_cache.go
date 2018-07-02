package nut

import (
	"container/list"
	"github.com/nut/common"
)

type lruCache struct {
	ks  *list.List
	kvs map[interface{}]interface{}
	max int
}

var lruDv = []byte{0}

func NewLRUCache(max int) *lruCache {
	return &lruCache{
		ks:  list.New(),
		kvs: make(map[interface{}]interface{}),
		max: max,
	}
}

func (lm *lruCache) Len() int {
	if lm.ks != nil && lm.kvs != nil {
		return lm.ks.Len()
	} else {
		return 0
	}
}

func (lm *lruCache) Contains(value interface{}) bool {
	return common.IsNotNil(lm.kvs[value])
}

func (lm *lruCache) Add(value interface{}) {
	lm.AddIfAbsent(value)
}

func (lm *lruCache) AddIfAbsent(value interface{}) interface{} {
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
		lm.kvs[value] = lruDv
		return nil
	}
}

func (lm *lruCache) Remove(value interface{}) bool {
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

func (lm *lruCache) Foreach(consumer func(...interface{})) {
	for e := lm.ks.Front(); e != nil; e = e.Next() {
		consumer(e.Value)
	}
}

func (lm *lruCache) ForeachBreak(bk func(...interface{}) bool, consumer func(...interface{})) interface{} {
	for value := range lm.kvs {
		if b := bk(value); b {
			return value
		}
		consumer(value)
	}
	return nil
}

func (lm *lruCache) Clear() {
	lm.ks.Init()
	lm.kvs = make(map[interface{}]interface{})
}
