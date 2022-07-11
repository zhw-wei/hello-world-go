
package main

import "fmt"

import "hello/hello"
import "hello/helloSql"

func main() {
	fmt.Println("hello world")

	fmt.Println("----------")

	hello.Main0_base()

	fmt.Println("----------")

	hello.Main2_base()

	fmt.Println("----------")

	helloSql.ConnectMysql()

}