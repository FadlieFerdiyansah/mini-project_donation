package models

import (
	"mini_project/config"
	"mini_project/helpers"
	"time"
)

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CheckLogin(email, password string) (bool, error) {
	var obj User
	con := config.InitDB()
	err := con.Where("email = ?", email).First(&obj).Error
	if err != nil {
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, obj.Password)
	if !match {
		return false, err
	}

	return true, nil
}
