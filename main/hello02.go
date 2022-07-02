package main

import "fmt"

func main3(){

	fmt.Println("hello world")

	var book1 Books
	book1.title = "golang"
	book1.author = "zhw"
	book1.subject = "demo"
	book1.bookId = 1

	fmt.Println(book1)
}

//定义结构体
type Books struct{
	title string
	author string
	subject string
	bookId int
}