package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

// 我们可以把一个方法关联在一个结构体上：
func (s *Saiyan) Super() {
	s.Power += 10000
}

/*
 将方法关联到 结构体上， 有点类型 js 的 bind() 方法。
 我们可以 认为 s 就是 js 里边的 this，
 这里的 s 代表的就是指针地址，不是指针副本是真正的地址
 通过  <instance>.Super(),  就可以操作 实例上的属性值了
*/
func main() {
	wukong := &Saiyan{"WuKong", 8800}
	wukong.Super()
	fmt.Println(wukong.Power) // => 18800
}
