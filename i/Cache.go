package i

type Cache interface {
	Len() int
	Put(interface{})
	Get(interface{})
	Remove(interface{})
}
