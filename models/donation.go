package models

import (
	"time"
)

type Donation struct {
	ID              uint64     `json:"id"`
	UserID          uint64     `json:"user_id"`
	ProgramID       uint64     `json:"program_id"`
	PaymentMethodId uint64     `json:"payment_method_id"`
	UniqueId        string     `json:"unique_id"`
	Amount          uint64     `json:"amount"`
	Status          string     `json:"status"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}
