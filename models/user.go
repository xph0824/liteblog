package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name   string `gorm:"unique_index"`
	Email  string `gorm:"unique_index"`
	Avatar string
	Pwd    string
	Role   int    `gorm:"default:1"` // 0 管理员 1正常用户
}