package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var db *gorm.DB

func init()  {
	var err error
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/goproject?charset=utf8");
	if err != nil {
		panic("连接数据库失败")
	}else {
		fmt.Println("数据库连接成功")
	}
}

func GetDB() *gorm.DB {
	return db
}