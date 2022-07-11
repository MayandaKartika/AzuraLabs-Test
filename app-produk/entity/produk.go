package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id     uint    `gorm:"primaryKey;autoIncrement:true"`
	Name   string  `gorm:"type:varchar(100)"`
	Price  int
	Rating float32
	Likes  int
}