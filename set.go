package nut

import (
	"github.com/nut/i"
)

// type Set interface {
// 	Len() int
// 	Contains(interface{}) bool
// 	Add(interface{}) bool
// 	Remove(interface{}) bool
// 	ContainsAll(Set) bool
// 	AddAll(Set) bool
// 	RetainAll(Set) bool
// 	RemoveAll(Set) bool
// 	Clear()
//
// 	Iterator
// }
var defaultValue = []byte{0}

type SetImpl struct {
	kvs map[interface{}]interface{}
	nf  func(interface{}) bool
	eq  func(interface{}, interface{}) bool
}

func NewSet() *SetImpl {
	return &SetImpl{kvs: make(map[interface{}]interface{})}
}

func NewSet2(nf func(interface{}) bool) *SetImpl {
	return &SetImpl{
		kvs: make(map[interface{}]interface{}),
		nf:  nf,
	}
}

func (si *SetImpl) Len() int {
	if si.kvs != nil {
		return len(si.kvs)
	} else {
		return 0
	}
}

func (si *SetImpl) Contains(value interface{}) bool {
	if si.isNil(value) {
		panic("value nil")
	}
	return si.kvs[value] != nil
}

func (si *SetImpl) Add(value interface{}) bool {
	if si.isNil(value) {
		panic("value nil")
	}
	
	if !si.isNil(si.kvs[value]) {
		return false
	}
	si.kvs[value] = defaultValue
	return true
}

func (si *SetImpl) Remove(value interface{}) bool {
	if si.isNil(value) {
		panic("value nil")
	}
	
	if !si.Contains(value) {
		return false
	}
	delete(si.kvs, value)
	return true
}

func (si *SetImpl) Foreach(consumer func(...interface{})) {
	for key := range si.kvs {
		consumer(key)
	}
}

func (si *SetImpl) ForeachBreak(bk func(interface{}) bool, consumer func(...interface{})) interface{} {
	for key := range si.kvs {
		if b := bk(key); b {
			return key
		}
		consumer(key)
	}
	return nil
}

func (si *SetImpl) ContainsAll(set i.Set) bool {
	_ := set.ForeachBreak(func(value interface{}) bool {
		return si.kvs[value] == nil
	}, func(consumer ...interface{}) {
	
	})
	return true
}

func (si *SetImpl) isNil(value interface{}) bool {
	return value == nil || (si.nf != nil && si.nf(value))
}

func (si *SetImpl) isEq(value1, value2 interface{}) bool {
	if si.eq != nil {
		return si.eq(value1, value2)
	} else {
		return value1 == value2
	}
}
