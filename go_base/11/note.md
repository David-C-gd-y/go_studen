go 的构造器

构造体没有 构造器，但是，我们可以创建一个返回所期望的类型的实例函数，类似工厂函数


```go
type Saiyan struct {
    Name string
    Power int
}

func NewSaiyan(name string power int) Saiyan {
    return Saiyan {
        Name: name,
        Power: power,
    }
}
```