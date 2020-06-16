package jwt

import (
	"github.com/gin-gonic/gin"
	"go-xb-song/applications/app/pkg/util"
	"net/http"
	"time"
)


func Jwt() gin.HandlerFunc {

	return func(c *gin.Context) {

		code := 1

		// token 通过header传递
		token := c.GetHeader("token")
		if token == "" {
			// token 为空
			code = 0
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				// 出错
				code = 0
			} else if time.Now().Unix() > claims.ExpiresAt {
				// 超时
				code = 0
			}
		}
		if code != 1 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
