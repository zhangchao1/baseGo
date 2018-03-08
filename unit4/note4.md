## 数组和切片
可以使用for range结构用来切分数据
```
for ix, value := range slice1 {
	...
}
第一个返回值 ix 是数组或者切片的索引，第二个是在该索引位置的值；他们都是仅在 for 循环内部可见的局部变量。value 只是 slice1 某个索引位置的值的一个拷贝，不能用来修改 slice1 该索引位置的值。
```
### 数组
数组在声明的时候，必须制定数组的长度，因为go在编译的时候，会使用
数组的指针传递和值传递的区别可以用下边两个示例来体现
示例1：值传递
```
func valueArray(a [3]int) {
	fmt.Println(a)
	a = [3]int{40, 50, 60}
}
```
示例2：指针传递
```
func pointerArray(a *[3]int) {
	fmt.Println(a)
	*a = [3]int{50, 60, 70}
}
```
然后分别在底下进行调用，我们就会发现
```
var b = [3]int{10, 20, 30}
	valueArray(b)
	fmt.Println(b)
	pointerArray(&b)
	fmt.Println(b)
最后的输出结果是b数组是：[50 60 70]
两个例子说明按值传递不会影响调用方的数据，按指针进行传递，由于当改变参数的值，就会影响调用方的值，因为两个变量的地址是一致的
```
#### 数组的初始化
数组的初始化的时候，举例说明
```
var b = [5]bool{}
	fmt.Println(b)

	var d = [5]int64{}
	fmt.Println(d)

	var e = [5]string{}
	fmt.Println(e)
当数组某些下标是空的时候，bool类型用false填充，string用空字符串填充，int和float用0填充
```
#### 二维数组的初始化
```
var b = [3][5]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}
```

### 切片
切片是可变长度的数组,而且切片本身就是指针，是一个引用类型
len是切片的长度，cap是切片的长度+数组除切片之外的长度,切片一定满足下面的不等式
0 <= len(s) <= cap(s)。

new和make的区别
new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址：这种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体（参见第 10 章）；它相当于 &T{}。
make(T) 返回一个类型为 T 的初始值，它只适用于3种内建的引用类型：切片、map 和 channel
示例如下：
```
func diffNewAndMake() {
	var b *[]int = new([]int)
	*b = []int{1, 2, 3}
	fmt.Println(b)
	c := b
	*c = []int{6, 7, 89}
	fmt.Println(b, c)

	var b1 = make([]string, 10, 50)
	fmt.Println(b1)
	a1 := b1
	a1[6] = "m"
	fmt.Println(a1, b1)
}
b1和b的值都被c和a1改了，因为切片是引用的
```
切片重组
改变切片长度的过程称之为切片重组 reslicing，做法如下：slice1 = slice1[0:end]，其中 end 是新的末尾索引（即长度）。

切片append以后，cap的容量的判断
这样举例
```
var src = []int{40, 30, 20, 10, 0}
这个时候，src的长度是5，容量是5
app是新添加的切片
newSrc = append(src,app)
newSrc的长度是len(src) + len(app)
对于newSrc的容量的计算，我们可以通过源码来进行计算
1.当len(src)+len(app)<2*len(src),那么cap(newSrc)的长度是2*len(src)
2.当len(src)+len(app)>2*len(src),那么cap(newSrc)的长度是len(Src)+len(app)
最后cap(newSrc)要进行内存对齐了
```
对于内存对齐的问题，可以参考如下链接：
[参考slice的growSlice](https://golang.org/src/runtime/slice.go)

copy slice的时候
func copy(dst, src []T) int copy 方法将类型为 T 的切片从源地址 src 拷贝到目标地址 dst，覆盖 dst 的相关元素，并且返回拷贝的元素个数。源地址和目标地址可能会有重叠。拷贝个数是 src 和 dst 的长度最小值。如果 src 是字符串那么元素类型就是 byte。如果你还想继续使用 src，在拷贝结束后执行 src = dst。