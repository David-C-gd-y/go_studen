学习 如何导入 package

go 中很多原生 内置的函数， 但是当我们使用第三方工具包的时候，是显式声明导入 工具包的，

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("It's over", os.Args[1])
}

```

内置函数 len， 可以截取命令行启动参数， 通过 os.Args 截取到对应的结果
os.Args 默认值 第 0 个，是一个文件运行的绝对路径