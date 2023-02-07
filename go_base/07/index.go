package main

import (
	"fmt"
	"os"
)

func log1() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("It's over", os.Args[0], os.Args[1])
	// It's over C:\Users\Admin\AppData\Local\Temp\go-build3040228429\b001\exe\index.exe 9000

}

func log2() {
	s, sep := "", ""
	/*
		通过 range 语法 标识 range os.Args[n,m] 的返回结果是一个切片类型的结果 ，
		类似 slice 操作， [0:] 表示 从 0 开始 结束
		也可以显式声明 [start: end] 这样的格式
		执行 go run index.go 9000 1 2 3
		输出 C:\Users\Admin\AppData\Local\Temp\go-build2865564574\b001\exe\index.exe 9000 1 2 3
	*/
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func main() {

	log2()
}
