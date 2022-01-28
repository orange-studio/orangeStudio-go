package apis

import (
	"go-study-server/src/utils"

	"github.com/gin-gonic/gin"
)

func GetUUID64(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": utils.GetUUID64(),
	})
}

func GetUUID32(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": utils.GetUUID32(),
	})
}
