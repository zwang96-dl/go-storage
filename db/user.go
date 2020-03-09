package db

import (
	"fmt"

	mydb "github.com/nicemayi/go-storage/db/mysql"
)

func UserSignup(username string, passwd string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"INSERT IGNORE INTO tbl_user (`user_name`, `user_pwd`) VALUES (?, ?)",
	)

	if err != nil {
		fmt.Println("Failed to prepare, err: " + err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(username, passwd)
	if err != nil {
		fmt.Println("Failed to insert, err: " + err.Error())
		return false
	}

	rowsAffected, err := ret.RowsAffected()
	if err == nil && rowsAffected > 0 {
		return true
	}
	if err != nil {
		fmt.Println("Failed to insert, err: " + err.Error())
		return false
	}
	fmt.Println("Already have record for username " + username)
	return false
}

func UserSignin(username string, encpwd string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"SELECT * FROM tbl_user WHERE user_name=? LIMIT 1",
	)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	rows, err := stmt.Query(username)
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else if rows == nil {
		fmt.Println("Username not found " + username)
		return false
	}

	pRows := mydb.ParseRows(rows)
	if len(pRows) > 0 && string(pRows[0]["user_pwd"].([]byte)) == encpwd {
		return true
	}

	return false
}

func UpdateToken(username string, token string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"REPLACE INTO tbl_user_token (`user_name`, `user_token`) VALUES (?, ?)",
	)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	defer stmt.Close()

	_, err = stmt.Exec(username, token)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
