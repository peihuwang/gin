/**
 * Created by Goland.
 * User: TinyStar
 * Date: 2020/01/19
 * Time: 10:56
 */
package main

import (
	"gin/bootstrap"
	"gin/route"
	_ "github.com/go-sql-driver/mysql"
)


func main() {

	//启动配置文件
	bootstrap.Bootstrap()

	r := route.SetupRouter()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8090")

}
