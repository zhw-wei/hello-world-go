package helloSql

import "fmt"

func UpdateUser() {
	//执行SQL语句
	sql := "update user set username =? where user_id = ?"
	res, err := db.Exec(sql, "user002", 2)

	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}

	//查询影响的行数，判断修改插入成功
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed", err)
	}

	fmt.Println("update succ:", row)
}
