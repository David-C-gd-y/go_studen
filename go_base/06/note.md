打包第一个 go 应用程序

执行一次文件，看看运行环境是否正常
```bash
go run index.go 
```

打包一个 go 则执行
```bash
go run buld index.go
```
会的一个 index.exe 可执行文件 ，名为 index ，你可以执行该文件。如果是在 Linux / OSX 系统中，别忘了使用 ./ 前缀来执行，也就是输入 ./index 。

在开发中，你既可以使用 go run 也可以使用 go build 。但当你正式部署代码的时候，你应该部署通过 go build 产生的二进制文件并且执行它。



查看临时文件的位置
 临时文件是哪里来的， 我们执行 go run index.go 的时候，会产生一个可执行的临时文件，只不过当执行完毕以后 会被删除掉

 如何查看这个临时文件？

 添加 命令行参数 --work 会得到临时文件的路径信息

```bash
go run --work index.go
WORK=C:\Users\Admin\AppData\Local\Temp\go-build1567497230
```
