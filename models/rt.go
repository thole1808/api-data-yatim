package models

import "time"

type RT struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Number    string    `json:"number"`
	RWID      uint      `json:"rw_id"`
	CreatedBy *uint     `json:"created_by"` // nullable
	UpdatedBy *uint     `json:"updated_by"` // nullable
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
