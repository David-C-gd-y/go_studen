package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	wukong := new(Saiyan)
	wukong.Name = "悟空"
	wukong.Power = 1000

	wufan := &Saiyan{
		Name:  "悟饭",
		Power: 800,
	}
	// 上面两种方式 使用起来是等价的
	fmt.Println(wufan, wukong)

}
