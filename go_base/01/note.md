1. 使用go module模式新建项目
    使用go module模式新建项目时，我们需要通过go mod init 项目名命令对项目进行初始化，该命令会在项目根目录下生成go.mod文件。例如，我们使用hello作为我们第一个Go项目的名称，执行如下命令。
```bash
go mod init hello
```

2. 运行 hello world

创建工作目录 ./

文件名: test.go，代码如下：

```go
package main

import "fmt"

func main() {
   fmt.Println("Hello, World!")
}
```
运行 命令行

```go
go run test.go

---> Hello, World!
```