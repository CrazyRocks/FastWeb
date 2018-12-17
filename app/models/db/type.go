package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"../../../config"
)

func Db() *gorm.DB {
	linkStr := config.DB_USER +":" + config.DB_PASS + "@(" + config.DB_HOST + ":" + config.DB_PORT + ")" + "/" + config.DB_NAME + "?charset=utf8&parseTime=True&loc=Local"
	db, error := gorm.Open("mysql", linkStr)
	if error != nil {
		fmt.Println("数据库发生错误")
		fmt.Println(error.Error())
	}
	return db
}
