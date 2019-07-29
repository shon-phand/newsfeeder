package handler

import (
	"github.com/gin-gonic/gin"
)

// PingGet function
func PingGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Hello ": "world",
		})
	}
}
