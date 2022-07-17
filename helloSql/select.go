package helloSql

import "fmt"

func SelectUser() {
	var users []Users
	sql := "select user_id, username,sex,email from user where user_id=? "
	err := db.Select(&users, sql, 2)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("select succ:", users)
}
