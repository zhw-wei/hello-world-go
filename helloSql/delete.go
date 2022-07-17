package helloSql

import "fmt"

func DeleteUser() {
	sql := "delete from user where user_id=?"

	res, err := db.Exec(sql, 2)
	if err != nil {
		fmt.Println("exce failed,", err)
		return
	}

	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("row failed, ", err)
	}
	fmt.Println("delete succ: ", row)
}
