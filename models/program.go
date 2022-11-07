package models

import (
	"time"
)

type Program struct {
	ID             uint64     `json:"id"`
	UserId         uint64     `json:"user_id"`
	CategoryId     uint64     `json:"category_id"`
	Title          string     `json:"title"`
	Body           string     `json:"body"`
	ExpectedAmount int64      `json:"expected_amount"`
	ExpiredDate    string     `json:"expired_date"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}
