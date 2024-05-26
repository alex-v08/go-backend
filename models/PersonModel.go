package models

import (
	"time"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Id 	  uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time

}

func (Person) TableName() string {
	return "person"
}

