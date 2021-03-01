package model

import "github.com/jinzhu/gorm"

type DeliverServer struct {
	gorm.Model
	Userid int	`gorm:"index"`
	Serverid int `gorm:"index"`
	Status int
}
