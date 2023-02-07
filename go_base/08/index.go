package main

// func log(message string) {
// }

// func add(a int, b int) int {
// 	return 2
// }

// 准确的变量名可以传达函数返回值的含义。尤其在返回值的类型都相同时，就像下面这样：
func power(name string) (int, bool, int) {

	if name == "goku" {
	}

	return 1, false, 2
}

/*
	空标识符 ， 即_（也就是下划线）
	为什么会有空标识符出现，因为在 go 里边不允许 未使用的变量声明存在，这样会导致编译错误
	有时候，你仅仅关注其中一个返回值。这个情况下，你可以将其他的返回值赋值给空白符_：
*/

func main() {

	_, exists, _ := power("goku")
	if exists == false {
		// 处理错误情况
	}

}
