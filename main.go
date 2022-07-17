package main

import (
	"fmt"
	"hello/hello"
	"hello/helloSql"
)

func main() {
	fmt.Println("hello world")

	fmt.Println("----------")

	hello.Main0_base()

	fmt.Println("----------")

	hello.Main2_base()

	fmt.Println("----------")

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
