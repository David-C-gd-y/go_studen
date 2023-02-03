变量 和 声明

默认情况下，Go 会为变量分配默认值。Integers 的默认值是 0，booleans 默认值是 false，strings 默认值是 ""

```go
	var power int // 没有赋值
	fmt.Printf("默认是 %d\n", power) // => 默认是  0
```

可以自动推导的 局部变量声明方式

```go
    power := 1
    fmt.Printf("%d\n", power) //  => 1
// 使用 := 方式声明的变量， 无法在全局作用域 使用
```

声明的变量如果未使用，编译会出错

```go
	name, power3 := "Goku", 9000
	fmt.Printf("default power is %d\n", power3) // => name declared and not used
```
这将不会通过编译，因为 name 是一个被声明但是未被使用的变量，就像 import 的包未被使用时，也将会导致编译失