// models/jenis_pembayaran.go
package models

type JenisPembayaran struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
}

func (JenisPembayaran) TableName() string {
	return "jenis_pembayaran"
}
