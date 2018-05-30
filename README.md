### NUT 是GO语言工具库

NUT 是GO语言工具库，提供一些常用数据结构的封装。

### LinkedMap

1, 创建 `lm := NewLinkedMap()`

2, 赋值

`Put`

`PutIfAbsent`

> 1 如果不存，放置，返回nil
> 2 否则，返回原值

`ComputeIfAbsent`

> 1 如果不存在，跟据func产生新值
> 1.1 如果新值不为nil，放置，返回nil
> 1 1.2 否则返回原值

`ComputeIfPresent`

> 1 如果不存在，不做任何处理，返回nil
> 2 否则根据func生成新值
> 2.1 如果新值为nil，删除key，返回nil
> 2.2 否则，放置，返回新值