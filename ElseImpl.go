package nut

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

func NewTuple() *TupleImpl {
	return &TupleImpl{}
}
