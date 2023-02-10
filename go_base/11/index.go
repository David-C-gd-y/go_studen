package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

// 工厂函数不必返回指针
// func NewSaiyan(name string, power int) *Saiyan {
// 	return &Saiyan{
// 		Name:  name,
// 		Power: power,
// 	}
// }

// 正确示例
func NewSaiyan(name string, power int) Saiyan {
	return Saiyan{
		Name:  name,
		Power: power,
	}
}

func main() {

	fmt.Println(NewSaiyan("wukong", 8888))
}
