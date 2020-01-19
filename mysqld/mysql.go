/**
 * Created by Goland.
 * User: TinyStar
 * Date: 2019/11/11
 * Time: 10:56
 */
package mysqld

import (
	"database/sql"
	"fmt"
	"gin/config"
	_"log"
)

var Db *sql.DB

func DbConnect(){
	var err error
	username :=config.ConfSet.Database.UserName
	password :=config.ConfSet.Database.Password
	host :=config.ConfSet.Database.Host
	port :=fmt.Sprintf("%d",config.ConfSet.Database.Port)

	dataSourceName :=username+":"+password+"@tcp("+host+":"+port+")/gogin?parseTime=true"

	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil{
		//log.Fatalln(err)
	}
	//defer Db.Close()


	Db.SetMaxIdleConns(20)
	Db.SetMaxOpenConns(20)

	if err := Db.Ping(); err != nil{
		//log.Fatalln(err)
	}
}
