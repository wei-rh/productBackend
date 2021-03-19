package model

import "github.com/jinzhu/gorm"

type Server struct {
	gorm.Model
	Userid int	`gorm:"index"`
	Tel string
	Name string
}
