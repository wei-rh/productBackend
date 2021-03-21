package model

import "github.com/jinzhu/gorm"

type BuyServer struct {
	gorm.Model
	Buyid int	`gorm:"index"`
	Userid int	`gorm:"index"`
	Serverid int `gorm:"index"`
	Status int
}
