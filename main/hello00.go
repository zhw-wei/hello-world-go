//文件目录，和java类似
package main

//导入依赖包
import "fmt"

//全局变量
var t int = 10

//go run hello.go
func main(){
	fmt.Println("hello world!")

	//变量赋值，支持类型推断
	//局部变量
	var a = "hello"
	var b string = "world"
	var c bool
	fmt.Println(a, b, c)

	//获取a的地址
	fmt.Println(&a)

	//函数调用
	fmt.Println(max(1,2))

	//函数返回多个值
	fmt.Println(swap("abc", "def"))

	//重新赋值全局变量
	t := 11
	fmt.Println(t)
}

func max(num1 int, num2 int) int {
	//此处如果没有分号，下面的if语句需要缩进
	var result int;

	if(num1 > num2) {
		result = num1
	} else {
		result = num2
	}

	return result
}

func swap(x string, y string) (string, string){
	return y, x
}