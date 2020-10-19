# 基础

## Go语言的基础特征

```text
1、 自动立即回收
2、 更丰富的内置类型
3、 函数多返回类型
4、 错误处理
5、 匿名函数和闭包
6、 类型和接口
7、 并发编程
8、 反射
9、 语言交互性
```

## Golang文件名

所有的Go源码文件都是以".go"结尾

## Go语言特征

1、Go的函数、变量、常量、自定义类型、包(`package`) 的命名方式

```text
1、首字符可以是任意的Unicode字符或者下划线
2、剩余字符可以是Unicode字符、下划线、数字
3、字符长度不限
```

2、Go只有25个关键字

``` go
break     default      func    interface  select
case      defer        go      map        struct
chan      else         goto    package    switch
const     fallthrough  if      range      type
continue  for          import  return     var
```

3、Go还有37个保留字

```go
Constants:   true  false  iota nil
Types:  int int8 int16 int32 int64
        uint uint8 uint16 uint32 uint64 uintptr
        float32 float64 complex128 complex64
        bool byte rune string error
Functions:  make len cap new append copy close delete
            complex real imag
            panic recover
```

4、可见性

```text
1、声明在函数内部，是函数的本地值，类似Private
2、声明在函数外部，是对当前包可见(包内所有.go文件都可见)的全局值，类似于protect
3、声明在函数外部且首字母大写是所有包可见的全局值，类似于public
```

## 内置的类型和函数

### 内置值类型

#### 1、值类型

```go
bool
int32,int64,int8,int16,int32,int64
uint32,uint64,uint8(byte),uint16,uint32,uint64
float32,float64
string
complex64,complex128
array
```

#### 2、引用类型

```go
slice   -- 序列数组(最常用)
map     -- 映射
chan    -- 管道
```

### 内置函数

Go语言拥有一些不需要进行导入操作就可以使用的内置函数。它们有时可以针对不同的类型进行操作，例如：len、cap和append，或必须使用系统级的操作，例如：panic。因此，它们需要直接获得编译器的支持

```go
append             -- 用来追加元素到数组、slice中，返回修改后的数组、slice
close              -- 主要用来关闭channel
delete             -- 从map中删除key对应的value
panic              -- 停止常规的goroutine (panic和recover: 用来做错误处理)
recover            -- 允许程序定义goroutine的panic动作
imag               -- 返回complex的实部 (complex、real imag: 用于创建和操作复数)
real               -- 返回complex的虚部
make               -- 用来分配内存，返回Type本身 (只能应用slice,map,channel)
new                -- 用来分配内存，主要用来分配值类型，比如int、struct。返回指向Type的指针
cap                -- capacity是容量的意思，用来返回某个类型的最大容量 (只能用于切片和map)
copy               -- 用来复制和连接slice,返回复制的数目
len                -- 来求长度，比如string、array、slice、map、channel，返回长度
print、println     -- 底层打印函数，在部署环境建议使用 fmt 包
```

### 内置接口error

```go
type error interface { // 只要实现Error函数，返回值为String的都实现了err接口
    Error() String
}
```

## int函数和main函数

### `init` 函数

go语言中`init`用于包`(package)`的初始化，该函数是go语言的一个重要特性。

``` text
1、init函数是用于程序执行前做包初始化的函数，比如初始化包里面的变量等
2、每个包可以拥有多个init函数
3、包的每个源文件也可以拥有多个init函数
4、同一个包中多个init函数的执行顺序go语言没有明确的定义 (说明)
5、不同的init的函数安装包导入的依赖的关系决定该初始化函数的执行顺序
6、init函数不能被其他函数调用，而是在main函数执行之前，自动被调用
```

### `main` 函数

```text
Go语言程序的默认入口函数(主函数)：func main()
函数体用{}一对括号包裹

func main(){
    // 函数体
}
```

`init`函数和`main`函数的异同

``` text
相同点：
    两个函数在定义时不能有任何的参数和返回值，且Go程序自动调用。

不同点：
    init可以应用于任意包中，且可以重复定义多个。
    main函数只可以用于main包中，且只能定义一个。
```
