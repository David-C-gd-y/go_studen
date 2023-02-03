package main

import (
	"fmt"
)

func usePower() int {
	return 10000
}

func main() {
	var power int // 没有赋值 默认是 0
	fmt.Printf("未赋值， %d\n", power)
	power = 9000
	fmt.Printf("赋值过了， %d\n", power)
	// 自动推导类型的 局部变量声明方式， 不能再全局作用域工作只能工作在局部作用域
	// 比较合适的使用场景是 字面量声明， 以及函数返回值
	power1 := 9000

	fmt.Printf("%d\n", power1)

	power2 := usePower()
	fmt.Printf("%d\n", power2)

	// := 支持连续 赋值， 用逗号做分割
	// 如果 用变量未必使用会编译出错
	name, power3 := "Goku", 9000
	fmt.Printf("%s's power is over %d\n", name, power3)
	// fmt.Printf("default power is %d\n", power3) // => name declared and not used

}
