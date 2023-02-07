学习函数声明

函数是可以返回多个值的。让我们看三个函数：一个没有返回值，一个有一个返回值，一个有两个返回值。

```go

func log(message string) {

}


func add(a int, b int) int {

	return 2
}

func power(name string) (int, bool, int) {

	if name == "goku" {
	}

	return 1, false, 2
}

func main() {

	_, exists, _ := power("goku")
	if exists == false {
		// 处理错误情况
	}
}

```