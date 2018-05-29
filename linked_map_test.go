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
	
	// 是否存在
	exit := lm.ContainsKey("a")
	fmt.Println("ContainsKey", exit)
	
	// 1 如果不存，放置，返回nil
	// 2 否则，返回原值
	value := lm.PutIfAbsent("a", "a2")
	fmt.Println("PutIfAbsent", value)
	
	// 1 如果不存在，跟据func产生新值
	// 1.1 如果新值不为nil，放置，返回nil
	// 1 1.2 否则返回原值
	value = lm.ComputeIfPresent("a", func(key, value interface{}) interface{} {
		return value.(string) + "1"
	})
	fmt.Println("ComputeIfPresent", value)
	
	// 1 如果不存在，不做任何处理，返回nil
	// 2 否则根据func生成新值
	// 2.1 如果新值为nil，删除key，返回nil
	// 2.2 否则，放置，返回新值
	value = lm.ComputeIfAbsent("b", func(key interface{}) interface{} {
		return "b2"
	})
	fmt.Println("ComputeIfAbsent", value)
	
	// 循环
	fmt.Println("Foreach kv")
	lm.Foreach(func(i ...interface{}) {
		fmt.Println(i[0], i[1])
	})
	
	// 循环且根据条件中断，返回中断时的key
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
