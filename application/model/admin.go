package model

import "github.com/kainonly/gin-extra/datatype"

type Admin struct {
	ID         uint64
	Username   string
	Password   string
	Role       datatype.JSONArray `gorm:"type:json"`
	Call       string
	Email      string
	Phone      string
	Avatar     string
	Status     bool
	CreateTime uint64
	UpdateTime uint64
}
