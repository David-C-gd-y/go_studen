package main

import "fmt"

type Saiyan struct {
	Name  string
	Age   int
	Power int
}

func editSaiyanBody(s Saiyan) {
	s.Power += 10000
}

func editSaiyanBodyFromPointer(s *Saiyan) {
	s.Power += 10000
}

func main() {

	saiyan1 := Saiyan{
		Name:  "the saiyan",
		Age:   10,
		Power: 9000,
	}

	// 如果 不显式 声明传递的参数 是指针的话
	// go 会直接 copy 一个新的 对象传递进去
	editSaiyanBody(saiyan1)

	fmt.Println(saiyan1.Power)

	// 在 struct 实例化的时候，
	// 加上 & 前缀，标记这个变量保存的是地址类型
	// 在 函数 参数 传递的时候 在形参前面 添加 * 标记这个类型是地址类型
	saiyan2 := &Saiyan{
		Power: 8888,
	}

	editSaiyanBodyFromPointer(saiyan2)

	fmt.Println(saiyan2)

	// 很显然 Saiyan 和 *Saiyan 是有关系的，
	// 虽然它们结构方面保持一致，却是两种不同的类型，
	// 没有标记 * 的属于 指针副本，标记了 * 属于真正的指针
	// 为什么要这么做，理由是因为
	/*
		复制一个指针的成本比 复制一个结构的成本要低很多
		指针的真正价值在于能够分享 它所指向的值
	*/

}
