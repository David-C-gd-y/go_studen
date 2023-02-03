package main

import "fmt"

/*
 go 语言 定义的数据类型


*/
func main() {
	/*
	 布尔型

	 布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。
	*/
	var defindForBooleanTrue bool = true
	var defindForBooleanFalse bool = false
	fmt.Println(defindForBooleanFalse, defindForBooleanTrue)
	/*
	 数字型

	 int 整数型

	 folat 浮点型
	*/

	var defindForNumberInt int = 1
	var defindForNumberFolat32 float32 = 1.111
	var defindForNumberFolat64 float64 = 1.321321

	fmt.Println(defindForNumberInt, defindForNumberFolat32, defindForNumberFolat64)

	/*
	 字符串类型
	 Go语言的字符串的字节使用UTF-8编码标识Unicode文本。
	*/

	var defindForString = "string"

	fmt.Println(defindForString)

	/*
		派生类型
		(a) 指针类型（Pointer）
		(b) 数组类型
		(c) 结构体类型(struct)
		(d) 联合体类型 (union)
		(e) 函数类型
		(f) 切片类型
		(g) 接口类型（interface）
		(h) Map 类型
		(i) Channel 类型
	*/

}
