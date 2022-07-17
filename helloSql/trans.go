package helloSql

import "fmt"

// TransUser 事务测试
func TransUser() {
	//开启事务
	conn, err := db.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return
	}

	//执行插入语句
	r, err := conn.Exec("insert into user(username, sex, email)values(?, ?, ?)", "user01", "man", "usre01@163.com")

	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback() //出现异常，进行回滚操作
		return
	}

	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)

	r, err = conn.Exec("insert into user(username, sex, email)values(?, ?, ?)", "user02", "man", "user02@163.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	id, err = r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)

	conn.Commit()
}
