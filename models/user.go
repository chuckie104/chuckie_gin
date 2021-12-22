package models

import "time"

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Password string `json:"password"`
	Birthday *time.Time `json:"birthday" gorm:"type:date"`
}

func (User) TableName() string {
	return "user"
}