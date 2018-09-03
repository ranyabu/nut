package iterf

type Set interface {
	Contains(interface{}) bool
	Add(interface{}) bool
	Remove(interface{}) bool
	ContainsAll(Set) bool
	AddAll(Set) bool
	RetainAll(Set) bool
	RemoveAll(Set) bool

	Collection
	Iterator
}
