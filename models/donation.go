package models

import "gorm.io/gorm"

type Donation struct {
	gorm.Model
	ID     uint64 `gorm:"primaryKey;autoIncrement:true"`
	UserID uint64
	User   User
}
