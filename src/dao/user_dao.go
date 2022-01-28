package dao

import (
	"go-study-server/src/config"
	"go-study-server/src/entity"
	"log"
	"strconv"
)

func GetUser(id int) entity.SysUser {
	var sysUser entity.SysUser
	e := config.DB.QueryRow("SELECT id, username, userpwd, nickname, status FROM sys_user where id = "+strconv.Itoa(id)).Scan(&sysUser.Id, &sysUser.Username, &sysUser.Userpwd, &sysUser.Nickname, &sysUser.Status)
	if e != nil {
		log.Print(e)
	}
	return sysUser
}
