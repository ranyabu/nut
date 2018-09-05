package iterf

type Cache interface {
	Contains(interface{}) bool
	Add(interface{}, interface{})
	AddIfAbsent(interface{}, interface{}) interface{}
	Remove(interface{}) bool
	
	Collection
	Iterator
}
