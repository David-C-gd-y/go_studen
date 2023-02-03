1. https://golang.google.cn/dl/ 下载 windows 版本 .msi ，下载完成一路点 next 即可

2. 配置 GO PROXY
```bash
go env -w GOPROXY=https://goproxy.cn,direct
```
3. GO MOD设置
```bash
go env -w GO111MODULE=on
```