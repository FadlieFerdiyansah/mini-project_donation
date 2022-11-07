package models

type Category struct {
	ID   uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name string
}
