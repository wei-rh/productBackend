package model

import "github.com/jinzhu/gorm"

type BuyOrder struct {
	gorm.Model
	Userid int	`gorm:"index"`
	SendAddress string
	GetAddress string
	Remarks string
	Context string
	Time string
	Baojia int64


	BuyServerID int `gorm:"index"`
	BuyServer BuyServer
}
