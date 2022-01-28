package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func initDBConfig() {
	log.Print("start init db config...")
	DB, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/orange_studio")
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		log.Fatal("open database fail", err)
		return
	}
	log.Print("db connected...")
}

func Init() {
	initDBConfig()
}
