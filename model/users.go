package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Openid string	`gorm:"index"`
	Tel string	`gorm:"index"`
}
