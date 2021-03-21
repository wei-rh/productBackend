package model

import "github.com/jinzhu/gorm"

type DeliverServer struct {
	gorm.Model
	Deliverid int	`gorm:"index"`
	Userid int	`gorm:"index"`
	Serverid int `gorm:"index"`
	Status int
}
