package models

import "time"

type RT struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Number    string    `json:"number"`
	RWID      uint      `json:"rw_id"`
	CreatedBy *string   `json:"created_by"`
	UpdatedBy *string   `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (RT) TableName() string {
	return "rt"
}
