package m

import (
	"testing"
	"fmt"
)

func TestLinkedMap(t *testing.T) {
	lm := NewLinkedMap()
	
	lm.Put("a", "a1")
	key := lm.ContainsKey("a")
	fmt.Println(key)
	lm.Put("b", "b1")
	fmt.Println(lm.PutIfAbsent("a", "a2"))
	
	fmt.Println(lm.ComputeIfPresent("a", func(key, value interface{}) interface{} {
		return value.(string) + "1"
	}))
	fmt.Println(lm.ComputeIfAbsent("b", func(key interface{}) interface{} {
		return "b2"
	}))
	
	lm.Foreach(func(i ...interface{}) {
		fmt.Println(i[0], i[1])
	})
	
	v := lm.ForeachBreak(func(i interface{}) bool {
		if i == "b" {
			return true
		}
		return false
	}, func(i ...interface{}) {
		fmt.Println(i[0])
	})
	
	fmt.Println(v)
	
}
