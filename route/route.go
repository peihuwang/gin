/**
 * Created by Goland.
 * User: liuyongs
 * Date: 2019/11/11
 * Time: 10:56
 */
package route

import (
	"gin/controller"
	"gin/log"
	"gin/middlwware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {


	r := gin.Default()
	//日记文件
	r.Use(log.WriteLogToFile())


	//中间件 ===========>全局中间件
	//r.Use(middlwware.TokenVerify) //验证JWT  其他验证方法自行补充

	//路由组
	user:=r.Group("/user")
	user.POST("/login",controller.UserLogin)
	user.Use(middlwware.TokenVerify)
	{
		user.POST("/list",controller.UserList)
		user.POST("/add",controller.UserAdd)
		user.POST("/delete",controller.UserDelete)
	}

	return r
}