package i

type Iterator interface {
	Foreach(func(...interface{}))
	ForeachBreak(func(interface{}) (bool, interface{}), func(...interface{}))
}
