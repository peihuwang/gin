/**
 * Created by Goland.
 * User: TinyStar
 * Date: 2019/11/11
 * Time: 10:56
 */
package middlwware

import (
	"gin/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TokenVerify(c *gin.Context)  {
	bools := jwt.ValidateToken(c.PostForm("token"))

	if(bools){
		c.Next()
	}else{
		c.Abort()

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "token is invalid",
		})
		return
	}
}