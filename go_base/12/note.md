结构体的 字段可以是任意类型，包括 其他结构体

```go
package main

import "fmt"

type Saiyan struct {
	Name   string
	Power  int
	Fahter *Saiyan
}

func main() {
	wufan := &Saiyan{
		Name:  "wufan",
		Power: 1000,
		Fahter: &Saiyan{
			Name:   "wukong",
			Power:  5000,
			Fahter: nil,
		},
	}

	fmt.Println(wufan)
}

```