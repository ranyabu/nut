package nut

import (
	"github.com/nut/i"
)

var setDv = []byte{0}

type setImpl struct {
	kvs map[interface{}]interface{}
	nf  func(interface{}) bool
	eq  func(interface{}, interface{}) bool
}

func NewSet() *setImpl {
	return &setImpl{kvs: make(map[interface{}]interface{})}
}

func NewSet2(nf func(interface{}) bool) *setImpl {
	return &setImpl{
		kvs: make(map[interface{}]interface{}),
		nf:  nf,
	}
}

func (si *setImpl) Len() int {
	return len(si.kvs)
}

func (si *setImpl) Contains(value interface{}) bool {
	if si.isNil(value) {
		panic("value nil")
	}
	return si.kvs[value] != nil
}

func (si *setImpl) Add(value interface{}) bool {
	if si.isNil(value) {
		panic("value nil")
	}
	
	if !si.isNil(si.kvs[value]) {
		return false
	}
	si.kvs[value] = setDv
	return true
}

func (si *setImpl) Remove(value interface{}) bool {
	if si.isNil(value) {
		panic("value nil")
	}
	
	if !si.Contains(value) {
		return false
	}
	delete(si.kvs, value)
	return true
}


func (si *setImpl) ContainsAll(set i.Set) bool {
	set.ForeachBreak(func(value ...interface{}) bool {
		return si.isNil(si.kvs[value[0]])
	}, func(consumer ...interface{}) {
		// pass
	})
	return true
}

func (si *setImpl) AddAll(set i.Set) bool {
	set.Foreach(func(value ...interface{}) {
		si.Add(value[0])
	})
	return true
}

func (si *setImpl) RemoveAll(set i.Set) bool {
	set.Foreach(func(value ...interface{}) {
		si.Remove(value[0])
	})
	return true
}

func (si *setImpl) Clear() {
	si.kvs = make(map[interface{}]interface{})
}

func (si *setImpl) Foreach(consumer func(...interface{})) {
	for key := range si.kvs {
		consumer(key)
	}
}

func (si *setImpl) ForeachBreak(bk func(interface{}) bool, consumer func(...interface{})) interface{} {
	for key := range si.kvs {
		if b := bk(key); b {
			return key
		}
		consumer(key)
	}
	return nil
}

func (si *setImpl) isNil(value interface{}) bool {
	return value == nil || (si.nf != nil && si.nf(value))
}

func (si *setImpl) isEq(value1, value2 interface{}) bool {
	if si.eq != nil {
		return si.eq(value1, value2)
	} else {
		return value1 == value2
	}
}
