package m

type Set interface {
	Len() int
	Contains(interface{}) bool
	Add(interface{}) bool
	Remove(interface{}) bool
	ContainsAll(Set) bool
	AddAll(Set) bool
	RetainAll(Set) bool
	RemoveAll(Set) bool
	Clear()

	Iterator
}
