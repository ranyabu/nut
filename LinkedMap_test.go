package nut

import (
	"testing"
	"fmt"
)

func TestLinkedMap(t *testing.T) {
	lm := NewLinedMap()
	
	lm.Put("c", "c")
	r := lm.ComputeIfPresent("c", func(key, value interface{}) interface{} {
		return value.(string) + "1"
	})
	fmt.Println(r)
}
