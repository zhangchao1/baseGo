### Map

#### 类型
Map是引用类型
Map的初始化，使用make进行初始化
可以用切片作为map的值
示例如下：
```
mp1 := make(map[int][]int)
mp2 := make(map[int]*[]int)
```
判断Map中元素是否存在，val1, isPresent = map1[key1]
删除map的元素 delete(map1,key1)
示例如下：
判断key是否存在
```
func checkKey() {
	var b = make(map[string]int)
	b = map[string]int{"a": 1, "b": 2}
	value, ok := b["a"]
	fmt.Println(value, ok)
	value, ok = b["d"]
	fmt.Println(value, ok)
}
```
删除key
```
func deleteKey() {
	var b = make(map[string]int)
	b = map[string]int{"a": 1, "b": 2}
	fmt.Println(b)
	delete(b, "a")
	fmt.Println(b)
}
```
#### for range构建map
```
for key, value := range map1 {
	...
}
```
第一个返回值 key 是 map 中的 key 值，第二个返回值则是该 key 对应的 value 值；这两个都是仅 for 循环内部可见的局部变量。其中第一个返回值key值是一个可选元素。