package models

import "time"

type ProfilYayasan struct {
	ID              uint       `json:"id" gorm:"primaryKey"`
	NamaYayasan     string     `json:"nama_yayasan"`
	Alamat          string     `json:"alamat"`
	NoTelepon       string     `json:"no_telepon"`
	Email           string     `json:"email"`
	Website         string     `json:"website"`
	TahunBerdiri    string     `json:"tahun_berdiri"`
	NoAktaPendirian string     `json:"no_akta_pendirian"`
	Pimpinan        string     `json:"pimpinan"`
	Visi            string     `json:"visi"`
	Misi            string     `json:"misi"`
	Deskripsi       string     `json:"deskripsi"`
	Logo            string     `json:"logo"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at" gorm:"index"`
}

func (ProfilYayasan) TableName() string {
	return "profil_yayasan"
}
