package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"productBackend/model"
)

func main() {
	//数据库连接
	url := "dbShop:123456@tcp(localhost:3306)/shop?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", url)
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	//db.DropTable(&model.Server{},&model.TakeOrder{},
	//&model.TakeServer{},&model.DeliverOrder{},&model.DeliverServer{},&model.BuyOrder{},&model.BuyServer{})

	//数据库迁移
	//db.Set("gorm:table_options", "ENGINE=InnoDB").
	//	AutoMigrate(&model.OrderProduct{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Server{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.TakeOrder{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.TakeServer{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.DeliverOrder{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.DeliverServer{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.BuyOrder{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.BuyServer{})


	log.Println("Migration completed")

}