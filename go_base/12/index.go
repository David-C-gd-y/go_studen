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
