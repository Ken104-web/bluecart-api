package model

import (
	"gorm.io/gorm"
	"time"
)


type User struct{
	gorm.Model
	Name string
	Email string

	SearchHistories []SearchHistory `gorm:"foreignKey:UserID"`

}

type SearchHistory struct{
	gorm.Model

	Name string
	SearchDate time.Time `gorm:"autoCreateTime"` 

	UserID uint // foreignKey to user 



}


