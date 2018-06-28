package nut

type Tuple interface {
	GetKey() interface{}
	GetValue() interface{}
}

type tupleImpl struct {
	K interface{}
	V interface{}
}

func (ti *tupleImpl) getKey() interface{} {
	return ti.K
}

func (ti *tupleImpl) getValue() interface{} {
	return ti.V
}

func NewTuple(k, v interface{}) *tupleImpl {
	return &tupleImpl{K: k, V: v}
}