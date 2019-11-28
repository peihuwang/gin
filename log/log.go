/**
 * Created by Goland.
 * User: liuyongs
 * Date: 2019/11/11
 * Time: 10:56
 */
package log

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
	"gin/config"
)

func WriteLog()  {
	//gin 自带的log
	gin.DisableConsoleColor()


	filename :=config.GetConfigPath()+`gin\log\day.log`

	file, _ := os.Create(filename)

	gin.DefaultWriter = io.MultiWriter(file)
}


// 日志写入文件
func WriteLogToFile() gin.HandlerFunc {

	fileName := config.GetConfigPath()+`gin\log\`+time.Now().Format("2006-01-02")+`.log`
	
	_, e := os.Stat(fileName)
	
	if e !=nil{
		os.Create(fileName)
	}

	src, _ := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	logger := logrus.New()

	logger.Out = src

	logger.SetLevel(logrus.DebugLevel)


	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})

	return func(c *gin.Context) {
		//谁请求的可以自己封装加入
		/*jwt.ValidateToken(c.PostForm("token"))
		userId :=jwt.Pl.UserID*/

		userId :=0

		reqStartTime := time.Now()

		c.Next()

		reqEndTime := time.Now()

		latencyTime := reqEndTime.Sub(reqStartTime)

		method := c.Request.Method

		reqUri := c.Request.RequestURI

		statusCode := c.Writer.Status()

		ip := c.ClientIP()

		bytes, _ := json.Marshal(c.Request.PostForm)
		body:=string(bytes)

		logger.Infof("| %3d | %3d | %13v | %15s | %s | %s | %s|",
			userId,
			statusCode,
			latencyTime,
			ip,
			method,
			reqUri,
			body,
		)
	}
}
