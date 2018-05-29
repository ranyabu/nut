package cache

type Cache interface {
	Len() int
	Contains(interface{}) bool
	Add(interface{})
	AddIfAbsent(interface{}) interface{}
	Remove(interface{}) bool
	Clear()
	
	Iterator
}
