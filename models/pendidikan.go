package models

type Pendidikan struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Level string `json:"level"`
}
