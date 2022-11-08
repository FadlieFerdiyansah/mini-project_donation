package models

import "gorm.io/gorm"

type Donation struct {
	gorm.Model
	ID            uint64 `json:"id"`
	UserID        uint64 `json:"user_id"`
	ProgramID     uint64 `json:"program_id"`
	PaymentMethod uint64 `json:"payment_method"`
	UniqueId      uint64 `json:"unique_id"`
	Amount        uint64 `json:"amount"`
	Status        uint64 `json:"status"`
	CreatedAt     uint64 `json:"created_at"`
	UpdatedAt     uint64 `json:"updated_at"`
	User          User
}
