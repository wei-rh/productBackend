package model

import "github.com/jinzhu/gorm"

type DeliverOrder struct {
	gorm.Model
	Userid int	`gorm:"index"`
	SendAddress string
	GetAddress string
	Remarks string
	Goods string
	Weight string
	Time string
	Baojia int64

	DeliverServerID int `gorm:"index"`
	// 关联
	DeliverServer DeliverServer
}
