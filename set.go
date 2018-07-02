package nut

import (
	"github.com/nut/iterf"
	"github.com/nut/common"
)

type setImpl struct {
	kvs map[interface{}]struct{}
}

func NewSet() *setImpl {
	return &setImpl{kvs: make(map[interface{}]struct{})}
}

func (si *setImpl) Len() int {
	return len(si.kvs)
}

func (si *setImpl) Contains(value interface{}) bool {
	return common.IsNotNil(si.kvs[value])
}

func (si *setImpl) Add(value interface{}) bool {
	if si.Contains(value) {
		return false
	}
	si.kvs[value] = struct{}{}
	return true
}

func (si *setImpl) Remove(value interface{}) bool {
	if !si.Contains(value) {
		return false
	}
	delete(si.kvs, value)
	return true
}

func (si *setImpl) ContainsAll(set iterf.Set) bool {
	set.ForeachBreak(func(value ...interface{}) bool {
		return common.IsNil(si.kvs[value[0]])
	}, func(consumer ...interface{}) {
		// pass
	})
	return true
}

func (si *setImpl) AddAll(set iterf.Set) bool {
	set.Foreach(func(value ...interface{}) {
		si.Add(value[0])
	})
	return true
}

func (si *setImpl) RetainAll(set iterf.Set) bool {
	modified := false
	set.Foreach(func(value ...interface{}) {
		if !si.Contains(value[0]) {
			si.Remove(value[0])
			modified = true
		}
	})
	return modified
}

func (si *setImpl) RemoveAll(set iterf.Set) bool {
	set.Foreach(func(value ...interface{}) {
		si.Remove(value[0])
	})
	return true
}

func (si *setImpl) Clear() {
	si.kvs = make(map[interface{}]struct{})
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
