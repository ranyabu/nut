package nut

import (
	"testing"
	"fmt"
)

func TestLinkedMap(t *testing.T) {
	lm := NewLinkedMap()

	lm.Put("a", "")
	key := lm.ContainsKey("a")
	fmt.Println(key)
}
