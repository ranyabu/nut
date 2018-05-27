package nut

//type Set interface {
//	Len() int
//	Contains(interface{}) bool
//	Add(interface{}) bool
//	Remove(interface{}) bool
//	ContainsAll(Set) bool
//	AddAll(Set) bool
//	RetainAll(Set) bool
//	RemoveAll(Set) bool
//	Clear()
//
//	Iterator
//}
type SetImpl struct {
	kvs map[interface{}]interface{}
}

func (si *SetImpl) Len() int {
	return len(si.kvs)
}

func (si *SetImpl) Contains(value interface{}) bool {
	return si.kvs[value] != nil
}

//func (si *SetImpl)
