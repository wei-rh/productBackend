package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"productBackend/config"
	"productBackend/handler"
	"productBackend/model"
	"productBackend/router"
)

func main() {
	//配置读取
	config.InitConf()
	// 初始化 DB
	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	takeorder := []model.TakeOrder{}
	model.DB.Find(&takeorder)
	fmt.Println(takeorder[0].Model.CreatedAt)
	//初始化路由的信息
	r := router.RouterInit()
	r.Run(config.App.ServiceAddr)
}
// 初始化 DB
func InitDB() (*gorm.DB, error){
	// 初始化 db（gorm），后边会封装起来
	db, err := gorm.Open("mysql", config.App.DSN)
	if err != nil {
		return nil, err
	}
	// 存储在其他代码可以访问的位置
	handler.DB = db
	model.DB = db
	return db, nil
}