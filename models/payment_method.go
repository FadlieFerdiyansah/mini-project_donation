package models

type PaymentMethod struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name        string
	PaymentCode string
}
