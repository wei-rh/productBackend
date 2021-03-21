package model

import "github.com/jinzhu/gorm"

type TakeServer struct {
	gorm.Model
	Takeid int `gorm:"index"`
	Userid int	`gorm:"index"`
	Serverid int `gorm:"index"`
	Status int
}
