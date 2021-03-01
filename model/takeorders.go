package model

import "github.com/jinzhu/gorm"

type TakeOrder struct {
	gorm.Model
	Userid int	`gorm:"index"`
	SendAddress string
	GetAddress string
	Remarks string
	Goods string
	Weight string
	Time string
	Baojia int64

	TakeServerID int `gorm:"index"`
	TakeServer TakeServer //
}
