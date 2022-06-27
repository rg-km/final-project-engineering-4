package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-4/backend/utils"
)

func Auth(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"timestamp": time.Now().Format("Mon, Jan 2 2006 15:04:05"),
				"code":      http.StatusBadRequest,
				"message":   "Silahkan Login",
				"data":      nil,
			})
			c.Abort()
			return
		}

		tokenStr = strings.ReplaceAll(tokenStr, "Bearer ", "")

		claims, err := utils.ValidateToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"timestamp": time.Now().Format("Mon, Jan 2 2006 15:04:05"),
				"code":      http.StatusBadRequest,
				"message":   err.Error(),
				"data":      nil,
			})
			c.Abort()
			return
		}

		if role != "all" && claims.Role != role {
			c.JSON(http.StatusUnauthorized, gin.H{
				"timestamp": time.Now().Format("Mon, Jan 2 2006 15:04:05"),
				"code":      http.StatusBadRequest,
				"message":   "Anda tidak memiliki izin untuk mengakses rute ini",
				"data":      nil,
			})
			c.Abort()
			return
		}

		c.Set("Email", claims.Email)
		c.Set("Role", claims.Role)
		c.Next()
	}
}
