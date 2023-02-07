package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("It's over", os.Args[0], os.Args[1])
	// It's over C:\Users\Admin\AppData\Local\Temp\go-build3040228429\b001\exe\index.exe 9000
}
