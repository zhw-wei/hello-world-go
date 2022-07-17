package main

import (
	"fmt"
	"hello/hello"
	"hello/helloGorm"
	"hello/helloSql"
)

func main() {
	runGorm()
}

func runHello() {
	fmt.Println("hello world")

	fmt.Println("----------")

	hello.Main0_base()

	fmt.Println("----------")

	hello.Main2_base()

	fmt.Println("----------")
}

func runHelloSql() {
	helloSql.CreateConnect()

	helloSql.InsertUser()

	fmt.Println("----------")
	helloSql.SelectUser()

	fmt.Println("----------")
	helloSql.UpdateUser()

	fmt.Println("----------")
	helloSql.SelectUser()

	fmt.Println("----------")
	helloSql.DeleteUser()

	fmt.Println("----------")
	helloSql.TransUser()
}

func runGorm() {

	helloGorm.CreateConnect()

}
