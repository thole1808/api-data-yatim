package models

import "time"

type Mitra struct {
	ID              uint           `json:"id"`
	Nama            string         `json:"nama"`
	Logo            *string        `json:"logo,omitempty"`
	Deskripsi       *string        `json:"deskripsi,omitempty"`
	TanggalMulai    *time.Time     `json:"tanggal_mulai,omitempty"`
	TanggalSelesai  *time.Time     `json:"tanggal_selesai,omitempty"`
	KategoriMitraID *uint          `json:"kategori_mitra_id"`
	KategoriMitra   *KategoriMitra `json:"kategori_mitra,omitempty" gorm:"foreignKey:KategoriMitraID"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

func (Mitra) TableName() string {
	return "mitra"
}

type KategoriMitra struct {
	ID         uint   `json:"id"`
	Nama       string `json:"nama"`
	Deskripsi  string `json:"deskripsi"`
	Keterangan string `json:"keterangan"`
}

func (KategoriMitra) TableName() string {
	return "kategori_mitra"
}
