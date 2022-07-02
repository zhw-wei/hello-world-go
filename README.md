# hello-world-go
go demo

```shell
# 对项目进行初始化，该命令会在项目根目录中生成go.mod文件
go mod init hello 
```

## 关键字
1. break, default, func, interface, select
2. case, defer, go, map, struct
3. chan, else, goto, package, switch
4. const, fallthrough, if, range, type
5. continue, for, import, return, var

## 预定义标识符
1. append, bool, byte, cap, close, complex, complex64, complex128, uint16
2. copy, false, float32, float64, imag, int, int8, int16, uint32
3. int32, int64, iota, len, make, new, new, nil, panic, unit64
4. print, println, real, recover, string, true, uint, uint8, uintprt

## 数据类型
1. 布尔类型: var b bool = true
2. 数字类型: 
    1. uint8, uint16, uint32, uint64, int8, int16, int32, int64
    2. float32, float64, complex64, complex128
    3. byte(类似uint8), rune(类似int32), uint(32或64), int(与uint一样大小), uintpr(无符号整型，用于存放一个指针)
3. 字符串类型
4. 派生类型

## 切片
切片实际上就是动态数组