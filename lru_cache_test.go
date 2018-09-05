package nut

import (
	"testing"
	"fmt"
)

func TestLRUCache(t *testing.T) {
	c := NewLRUCache(3)
	c.Add("k1","v1")
	fmt.Println(c.AddIfAbsent("k1","v2"))
	fmt.Println(c.Contains("k1"))
	c.Remove("k1")
	c.Add("k2","v2")
	c.Foreach(func(i ...interface{}) {
		fmt.Println(i)
	})
}
