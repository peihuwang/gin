/**
 * Created by Goland.
 * User: liuyongs
 * Date: 2019/11/11
 * Time: 10:56
 */
package controller

import (
	"crypto/sha1"
	"encoding/hex"
	"gin/entities"
	"gin/jwt"
	"gin/model"
	"github.com/gin-gonic/gin"
	"io"
	_"log"
	"net/http"
)

func UserLogin(c *gin.Context)  {
	name :=c.PostForm("name")
	password :=c.PostForm("password")

	hash := sha1.New()
	io.WriteString(hash, password)
	hSum := hash.Sum(nil)

	hashPassword := hex.EncodeToString(hSum)

	row := model.UserLogin(name)

	var userInfo entities.UserInfo

	err := row.Scan(&userInfo.Id, &userInfo.Name, &userInfo.Password,&userInfo.IsDelete)

	if err !=nil{
		//log.Fatalln(err)
	}

	msg :="登录失败"
	if userInfo.Password ==hashPassword{
		msg = "登录成功"
	}
	token :=jwt.MakeToke(userInfo.Id)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
		"token":token,
	})
}


func UserList(c *gin.Context){
	list := model.UserList()

	user := make([]entities.UserInfo, 0)

	for list.Next(){
		var userInfo entities.UserInfo
		list.Scan(&userInfo.Id,&userInfo.Name,&userInfo.Password,&userInfo.IsDelete)
		user = append(user,userInfo)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
		"currentId":jwt.Pl.UserID,
	})
}


func UserAdd(c *gin.Context)  {
	name :=c.PostForm("name")
	password :=c.PostForm("password")

	hash := sha1.New()
	io.WriteString(hash, password)
	hSum := hash.Sum(nil)

	hashPassword := hex.EncodeToString(hSum)

	id := model.UserAdd(name, hashPassword)

	//通过id 直接写入
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func UserDelete(c *gin.Context)  {
	name :=c.PostForm("name")
	row := model.UserDelete(name)

	msg:="删除失败"
	if(row>0){
		msg = "删除成功"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}