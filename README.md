### NUT 是GO语言工具库

NUT 是GO语言工具库，提供一些常用数据结构的封装。

NUT 抽象了Map、Set、Iterator、Cache、Lock接口，且以此为基础实现了LinkedMap、Set、LRUCache、Lock。

已下为接口说明：

### Iterator

```go
type Iterator interface {
	Foreach(func(...interface{}))
	ForeachBreak(func(...interface{}) bool, func(...interface{})) interface{}
}
```

1 Foreach     定义了遍历操作，用户需要自定义消费遍历值函数   
2 ForeachBreak定义了遍历操作，但可以提前中断遍历，返回遍历值；用户需要自定义中止函数和消费遍历值函数   

### Map

```go
type Map interface {
	Len() int
	ContainsKey(interface{}) bool
	Get(interface{}) interface{}
	Put(interface{}, interface{}) interface{}
	PutIfAbsent(interface{}, interface{}) interface{}
	ComputeIfAbsent(interface{}, func(interface{}) interface{}) interface{}
	ComputeIfPresent(interface{}, func(interface{}, interface{}) interface{}) interface{}
	Remove(interface{}) interface{}
	PutAll(Map)
	Clear()

	Iterator
}
```

1  Len                返回当前Map大小   
2  ContainsKey        判断是否包含Key   
3  Get                获取对应的Value   
4  Put                赋值，如果没有，新增，否则覆盖，返回新值   
5  PutIfAbsent        赋值，如果没有，新增，返回Nil；否则，放回旧值   
6  ComputeIfAbsent    赋值，如果没有，则根据函数生成新值，当新值不为Nil赋值且返回新值，否则返回Nil；如果有，则返回Nil   
7  ComputeIfPresent   赋值，如果没有，返回Nil；如果有，则根据函数生成新值，当新值为Nil，删除Key，返回Nil，否则赋值，放回新值   
8  Remove             删除KeyValue   
9  PutAll             赋值全部   
10 Clear              重置   
11 Foreach            遍历   
12 ForeachBreak       遍历，中断时返回当前Key   

