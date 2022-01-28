package apis

import (
	"orangestudio-go/src/dao"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sysUser := dao.GetUser(id)
	c.JSON(200, gin.H{
		"message": sysUser,
	})
}
