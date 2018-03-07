### 函数

#### 按值传递

#### 按引用传递

#### 可变参数的函数
这个函数接受一个类似某个类型的 slice 的参数
示例如下：
```
func canChangeParameter(a ...string) {
	fmt.Println(a)
}
```
也可以传递一个空类型
示例如下：
```
func canChangeParameterInterface(a ...interface{}) {
	fmt.Println(a)
}
```
#### defer函数的使用
关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数（为什么要在返回之后才执行这些语句？因为 return 语句同样可以包含一些操作，而不是单纯地返回某个值）。
关键字 defer 的用法类似于面向对象编程语言 Java 和 C# 的 finally 语句块，它一般用于释放某些已分配的资源。

多个defer堆在一块执行的时候，会逆序执行，比如向如下的示例：
```
func runDeferOrder() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
```
defer可以做一些收尾的工作，比如说文件关闭，资源解锁，关闭数据库链接
#### 闭包函数
闭包函数:表示参数列表的第一对括号必须紧挨着关键字 func，因为匿名函数没有名称。花括号 {} 涵盖着函数体，最后的一对括号表示对该匿名函数的调用。
示例如下：
```
func runLambaFun() {
	lam := func(i string) { fmt.Println(i) }
	lam("hellow word")
}
```
闭包函数作为返回值的话，在执行过程中，闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量。
示例说明：
```
func add(a int) func(b int) int {
	return func(b int) int {
		return b + a
	}
}
调用的时候:
rlam := add(20),初始化闭包的时候，a的参数会被记录下来
所以，在下次调用的时候，
rlam(3)的结果是23，因为a的数值被记录下来了
```
使用闭包函数，可以用于调试，使用"log"和"runtime"两个包进行日志的调试
您可以使用 runtime 或 log 包中的特殊函数来实现这样的功能。包 runtime 中的函数 Caller() 提供了相应的信息，因此可以在需要的时候实现一个 where() 闭包函数来打印函数执行的位置：
示例如下：
```
where := func() {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s:%d", file, line)
}
```
内存缓存：当在进行大量的计算时，提升性能最直接有效的一种方式就是避免重复计算。通过在内存中缓存和重复利用相同计算的结果，称之为内存缓存。
特别好的例子就是斐波那契数列
```
const LIM = 41

var fibs [LIM]uint64
func fibonacci(n int) (res uint64) {
	// memoization: check if fibonacci(n) is already known in array:
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}
```





