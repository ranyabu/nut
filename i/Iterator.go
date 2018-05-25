package i

type Iterator interface {
	Foreach(func(...interface{}))
	hashNext() bool
	next() interface{}
}
