package models

type PaymentMethod struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	PaymentCode string `json:"payment_code"`
}
