package iterf

type Cache interface {
	Contains(interface{}) bool
	Add(interface{})
	AddIfAbsent(interface{}) interface{}
	Remove(interface{}) bool

	Collection
	Iterator
}
