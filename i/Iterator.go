package i

type Iterator interface {
	Foreach(func(...interface{}))
}