 结构体的声明 和 初始化

```go
type Saiyan struct {
    Name string
    Age int
    Power int    
}

```

这个看起来和 TypeSctipt 相似

```ts
type Saiyan = {
    Name: string
    Age: number
    Power: number
}
```
但是在 go 的环境中，结构体是参与代码实际执行的，和 ts 不一样的地方，是 ts 尽管出错了，编译照样通过，然后可以执行 js 代码。
