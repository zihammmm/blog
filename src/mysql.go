package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"
)

/*
使用mysql作为数据库
*/

type Conn struct {
	db *gorm.DB
}

var conn *Conn
var mu sync.Mutex

func GetConnection() *Conn {
	mu.Lock()
	defer mu.Unlock()

	if conn == nil {
		conn = &Conn{}
		conn.init()
	}
	return conn
}

func (conn *Conn) init() {
	if conn.db != nil {
		return
	}
	dsn := "blog:19971123.@tcp(118.178.89.205:3306)/blog?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		log.Fatalln(err)
		return
	}
	conn.db = db
}
