package nut

import (
	"testing"
	"fmt"
)

func TestLinkedMap(t *testing.T) {
	lm := NewLinkedMap()

	lm.Put("a", nil)
	key := lm.ContainsKey("a")
	fmt.Println(key)
}
