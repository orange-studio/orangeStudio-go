package apis

import (
	"go-study-server/src/utils"

	"github.com/gin-gonic/gin"
)

func Get_uuid_64(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": utils.Get_uuid_64(),
	})
}

func Get_uuid_32(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": utils.Get_uuid_32(),
	})
}
