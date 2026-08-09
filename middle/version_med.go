package middle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func VersionMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		version := c.Request.Header.Get("version")
		if version == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no version set header"})
			c.Abort()
			return
		}
		c.Next()

	}
}
