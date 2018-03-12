### 控制结构

#### if-else

if的用法：但要注意的是，使用简短方式 := 声明的变量的作用域只存在于 if 结构中

示例如下：
 ```
 if val := 10; val > max {
	// do something
}
 ```
我们在进行错误获取的时候，就可以用这中方法进行判断，代码结构会比较清晰

示例如下：
```
if err := file.Chmod(0664); err != nil {
	fmt.Println(err)
	return err
}
```

#### for
基于逻辑判断的for的使用

```
var i int = 5
for i>0{
	fmt.Println(i)
}
```

由于golang不存在while的使用，所以要用for实现无限循坏
```
for{

}
```
for range可以用于数组，字典等输出

#### continue和break
continue只能用于for循坏

一个 break 的作用范围为该语句出现后的最内部的结构，它可以被用于任何形式的 for 循环（计数器、条件判断等）。但在 switch 或 select 语句中（详见第 13 章），break 语句的作用结果是跳过整个代码块，执行后续的代码。

示例说明 break 的作用范围为该语句出现后的最内部的结构，不是在所有的地方
```
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if j > 5 {
				break
			}
			print(j)
		}
		print("  ")
	}
	结果是：012345  012345  012345
```

#### 标签 
break的跳转标签(label)必须放在循环语句for前面，并且在break label跳出循环不再执行for循环里的代码。
示例如下：
```
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				break LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
		fmt.Println(i)
	}
```

continue label会执行for循坏里面的代码

示例如下：
```
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
		fmt.Println(i)
	}
```
goto label的label(标签)既可以定义在for循环前面,也可以定义在for循环后面，当跳转到标签地方时，继续执行标签下面的代码。

label在for循环的里面：
```
func runGotoLabel1() {
	var i int = 0
loop:
	for i < 10 {
		if i == 4 {
			i++
			goto loop
		}
		fmt.Println(i)
		i++
	}
}
```

label在循坏的下面：
示例如下：
```
func runGotoLabel2() {
	for i := 0; i < 5; i++ {
		if i == 4 {
			goto loop
		}
		fmt.Println(i)
	}
loop:
	fmt.Println("hellow world")
}
```









