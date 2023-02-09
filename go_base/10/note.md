结构体上的函数
```go
type Saiyan struct {
  Name string
  Power int
}

func (s *Saiyan) Super() {
  s.Power += 10000
}
```

注意 不要在 main 函数中声明 Super ，错误原因暂时还没学到

这种关联的操作，其实就类型 javascript 的 bind 函数的操作。

```js
function fn1 () {
    this.value +=1;
}

class A {
    construct() {
        this.value = 0
    }
}
const A_Instance = new A();
const fn2 = fn1.bing(A_Instance)

console.log(fn2()) // => { value: 1 }
console.log(fn2()) // => { value: 2 }
console.log(fn2()) // => { value: 3 }
```

bind 函数将 this 的地址 改变成 A_Instance 这个实例的地址
 
```go

func (s *Saiyan) Super() {
  s.Power += 10000
}
/*
 s, 我们可以比喻为 bind 函数指定的 this，
 在函数执行的时候， s 始终指向 某个认定的 内存地址
*/
```