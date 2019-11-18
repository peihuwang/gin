/**
 * Created by Goland.
 * User: liuyongs
 * Date: 2019/11/11
 * Time: 10:56
 */
package model

import (
	"database/sql"
	"gin/mysqld"
	_"log"
)


func UserList() *sql.Rows{
	rows, err :=  mysqld.Db.Query("SELECT *  FROM user")

	if err != nil {
		//log.Fatalln(err)
	}

	return rows
}


func UserAdd(user,password string) int64 {
	res, err := mysqld.Db.Exec("INSERT into user (name,password) value (?,?)", user, password)

	if err !=nil{
	//	log.Fatalln(err)
	}

	id, err := res.LastInsertId()
	if err !=nil{
		//log.Fatalln(err)
	}

	return id
}

func UserLogin(name string)  *sql.Row{
	row:= mysqld.Db.QueryRow("SELECT *  FROM user where name=?",name)
	return row
}

func UserDelete(name string) int64 {
	stmt, err := mysqld.Db.Prepare("update user set is_delete = 1 where name=?")

	if err !=nil{
	//	log.Fatalln(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(name)

	if err !=nil{
		//log.Fatalln(err)
	}

	ra, err := result.RowsAffected()
	if err != nil {
		//log.Fatalln(err)
	}

	return ra
}
