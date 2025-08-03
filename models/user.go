package models

import "time"

type User struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Role          string    `json:"role"`
	RememberToken string    `json:"remember_token"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// type User struct {
// 	ID        uint      `json:"id" gorm:"primaryKey"`
// 	Username  string    `json:"username" gorm:"unique"`
// 	Password  string    `json:"password"`
// 	Role      string    `json:"role"`
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// }
