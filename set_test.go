package nut

import (
	"testing"
	"fmt"
)

func TestSet(t *testing.T) {
	s := NewSet()
	s.Add("key1")
	s.Add("key1")
	s.Foreach(func(i ...interface{}) {
		fmt.Println(i[0])
	})
	s.Remove("key1")
	fmt.Println(s.Contains("key1"))
}

