package nut

import (
	"testing"
	"fmt"
)

func TestLinkedMap(t *testing.T) {
	
	lm := NewLinkedMap()
	
	lm.Put("a", "a1")
	lm.Put("b", "b1")
	lm.Put("c", "c1")
	lm.Put("d", "d1")
	
	exit := lm.ContainsKey("a")
	fmt.Println("ContainsKey", exit)
	
	value := lm.PutIfAbsent("a", "a2")
	fmt.Println("PutIfAbsent", value)
	
	value = lm.ComputeIfPresent("a", func(key, value interface{}) interface{} {
		return value.(string) + "1"
	})
	fmt.Println("ComputeIfPresent", value)
	
	value = lm.ComputeIfAbsent("b", func(key interface{}) interface{} {
		return "b2"
	})
	fmt.Println("ComputeIfAbsent", value)
	
	fmt.Println("Foreach kv")
	lm.Foreach(func(i ...interface{}) {
		fmt.Println(i[0], i[1])
	})
	
	fmt.Println("ForeachBreak c kv")
	v := lm.ForeachBreak(func(i ...interface{}) bool {
		if i[0] == "c" {
			return true
		}
		return false
	}, func(i ...interface{}) {
		fmt.Println(i[0], i[1])
	})
	
	fmt.Println("break", v, lm.kvs[v])
}
