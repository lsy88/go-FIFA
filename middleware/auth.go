package middleware

import (
	"FIFA/controller/response"
	"FIFA/utils/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		if header == "" {
			response.ResponseError(c, response.ERR_Token_NotExist)
			c.Abort()
			return
		}
		auth := strings.Split(header, " ")
		if len(auth) != 2 || auth[0] != "Bearer" {
			response.ResponseError(c, response.ERR_Invalid_Token)
			c.Abort()
			return
		}
		//token := auth[1]就是解析到的token
		j := jwt.NewJWT()
		claims, err := j.ParseToken(auth[1])
		if err != nil {
			response.ResponseError(c, response.ERR_Invalid_Token)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
