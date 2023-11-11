package v1

import (
	"github.com/gin-gonic/gin"
	"go-gin/middleware/jwt"
	"net/http"
)

func Home(c *gin.Context) {
	name := c.Query("name")

	if name == "gin" {
		token, _ := jwt.CreateToken(name)
		c.JSON(200, gin.H{
			"code":  200,
			"msg":   "hello " + name,
			"token": token,
		})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 500,
		"msg":  "账号密码错误",
	})
}
func Refresh(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  "請登入",
		})
		c.Abort()
		return
	}
	newToken, err := jwt.RefreshToken(token)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"code":  200,
		"msg":   "refresh token",
		"token": newToken,
	})
}
func Parse(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  "請登入",
		})
		c.Abort()
		return
	}
	claims, err := jwt.ParseToken(token)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  "token 錯誤",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "parse token",
		"claims": claims,
	})
}
