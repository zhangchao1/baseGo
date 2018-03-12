### 结构体

#### 工厂方法创建结构体
示例如下：
```
func NewFac(id int64, name string) *fac {
	return &fac{id, name}
}
```
使用这种方式可以更好的实例化结构体，类似于工厂方法
如果强制外部这样使用的话，就必须要在内部将结构体小写，
示例如下：
```
type fac struct {
	id   int64
	name string
}
```
结构体使用new,map使用make

#### 结构体字段 可以备注一些字段的备注

示例如下：
```
type tagStruct struct {
	id   int64  "orm id"
	name string "string 123"
}
```
这一个可以使用反射包的方法获取
方法如下：
```
func refTag(tt tagStruct, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}
```
#### 结构内嵌
同样地结构体也是一种数据类型，所以它也可以作为一个匿名字段来使用，如同上面例子中那样。外层结构体通过 outer.in1 直接进入内层结构体的字段，内嵌结构体甚至可以来自其他包。内层结构体被简单的插入或者内嵌进外层结构体。这个简单的“继承”机制提供了一种方式，使得可以从另外一个或一些类型继承部分或全部实现。

命名冲突
当两个字段拥有相同的名字（可能是继承来的名字）时该怎么办呢？
外层名字会覆盖内层名字（但是两者的内存空间都保留），这提供了一种重载字段或方法的方式；
如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误（不使用没关系）。没有办法来解决这种问题引起的二义性，必须由程序员自己修正。

####值方法和指针方法的区别
值方法接收者的类型是非指针的数据类型
指针方法和值方法都可以在指针或非指针上被调用

#### 总结
在 Go 中，类型就是类（数据和关联的方法）。Go 不知道类似面向对象语言的类继承的概念。继承有两个好处：代码复用和多态。

在 Go 中，代码复用通过组合和委托实现，多态通过接口的使用来实现：有时这也叫 组件编程（Component Programming）。

许多开发者说相比于类继承，Go 的接口提供了更强大、却更简单的多态行为。