package nut

type Tuple interface {
	getKey() interface{}
	getValue() interface{}
}

type TupleImpl struct {
	K interface{}
	V interface{}
}

func (ti *TupleImpl) getKey() interface{} {
	return ti.K
}

func (ti *TupleImpl) getValue() interface{} {
	return ti.V
}

func NewTuple(k, v interface{}) *TupleImpl {
	return &TupleImpl{K: k, V: v}
}