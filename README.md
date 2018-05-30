### NUT 是GO语言工具库

NUT 是GO语言工具库，提供一些常用数据结构的封装。NUT 抽象了Map、Set、Iterator、Cache、Lock接口，且以此为基础实现了LinkedMap、Set、LRUCache、Lock。

已下为接口说明：

### Iterator

```go
type Iterator interface {
	Foreach(func(...interface{}))
	ForeachBreak(func(...interface{}) bool, func(...interface{})) interface{}
}
```

| 接口           | 说明                                       |
| ------------ | ---------------------------------------- |
| Foreach      | 定义了遍历操作，用户需要自定义消费遍历值函数                   |
| ForeachBreak | 定义了遍历操作，但可以提前中断遍历，返回遍历值；用户需要自定义中止函数和消费遍历值函数 |



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

| 接口               | 说明                                       |
| ---------------- | ---------------------------------------- |
| Len              | 返回当前Map大小                                |
| ContainsKey      | 判断是否包含Key                                |
| Get              | 获取对应的Value                               |
| Put              | 添加KV，如果没有，新增，否则覆盖，返回新值                   |
| PutIfAbsent      | 添加KV，如果没有，新增，返回Nil；否则，放回旧值               |
| ComputeIfAbsent  | 添加KV，如果没有，则根据函数生成新值，当新值不为Nil添加且返回新值，否则返回Nil；如果有，则返回Nil |
| ComputeIfPresent | 添加KV，如果没有，返回Nil；如果有，则根据函数生成新值，当新值为Nil，删除Key，返回Nil，否则添加，放回新值 |
| Remove           | 删除Key对应的Value                            |
| PutAll           | 添加KV全部                                   |
| Clear            | 重置                                       |
| Foreach          | 遍历                                       |
| ForeachBreak     | 遍历，中断时返回当前Key                            |



### Set

```go
type Set interface {
	Len() int
	Contains(interface{}) bool
	Add(interface{}) bool
	Remove(interface{}) bool
	ContainsAll(Set) bool
	AddAll(Set) bool
	RetainAll(Set) bool
	RemoveAll(Set) bool
	Clear()
	
	Iterator
}
```

| 接口           | 说明            |
| ------------ | ------------- |
| Len          | 返回当前Set大小     |
| Contains     | 是否包含          |
| Add          | 添加元素          |
| Remove       | 删除元素          |
| ContainsAll  | 是否包含全部指定元素集合  |
| AddAll       | 添加全部指定元素集合    |
| RetainAll    | 保留指定元素集合      |
| RemoveAll    | 删除指定的元素集合     |
| Clear        | 重置            |
| Foreach      | 遍历            |
| ForeachBreak | 遍历，中断时返回当前Key |



### Cache

```go
type Cache interface {
	Len() int
	Contains(interface{}) bool
	Add(interface{})
	AddIfAbsent(interface{}) interface{}
	Remove(interface{}) bool
	Clear()
	
	Iterator
}
```

| 接口           | 说明                      |
| ------------ | ----------------------- |
| Len          | 当前缓存大小                  |
| Contains     | 是否包含                    |
| Add          | 新增元素                    |
| AddIfAbsent  | 新增元素，如果存在，返回存在值，否则返回Nil |
| Remove       | 删除元素                    |
| Clear        | 重置                      |
| Foreach      | 遍历                      |
| ForeachBreak | 遍历，中断时返回当前Key           |



### 使用示例

比如LinekedMap

```go
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
```
