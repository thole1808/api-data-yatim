// package models

// import "time"

// type Activity struct {
// 	ID             uint            `json:"id" gorm:"primaryKey"`
// 	KategoriID     *uint           `json:"kategori_id"`
// 	Judul          string          `json:"judul"`
// 	Deskripsi      *string         `json:"deskripsi"`
// 	Lokasi         *string         `json:"lokasi"`
// 	TanggalMulai   *time.Time      `json:"tanggal_mulai"`
// 	TanggalSelesai *time.Time      `json:"tanggal_selesai"`
// 	Status         string          `json:"status" gorm:"default:draft"`
// 	Images         []ActivityImage `json:"images" gorm:"foreignKey:AktivitasID"`
// }

// func (Activity) TableName() string {
// 	return "aktivitas"
// }

// type ActivityImage struct {
// 	ID          uint   `json:"id" gorm:"primaryKey"`
// 	AktivitasID uint   `json:"aktivitas_id"`
// 	Path        string `json:"path"`
// 	Caption     string `json:"caption"`
// }

// func (ActivityImage) TableName() string {
// 	return "aktivitas_gambar"
// }

package models

import "time"

type Activity struct {
	ID             uint              `json:"id" gorm:"primaryKey"`
	KategoriID     *uint             `json:"kategori_id"`
	Kategori       *KategoriKegiatan `json:"kategori" gorm:"foreignKey:KategoriID"`
	Judul          string            `json:"judul"`
	Deskripsi      *string           `json:"deskripsi"`
	Lokasi         *string           `json:"lokasi"`
	TanggalMulai   *time.Time        `json:"tanggal_mulai"`
	TanggalSelesai *time.Time        `json:"tanggal_selesai"`
	Status         string            `json:"status" gorm:"default:draft"`
	Images         []ActivityImage   `json:"images" gorm:"foreignKey:AktivitasID"`
}

func (Activity) TableName() string {
	return "aktivitas"
}

type ActivityImage struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	AktivitasID uint   `json:"aktivitas_id"`
	Path        string `json:"path"`
	Caption     string `json:"caption"`
}

func (ActivityImage) TableName() string {
	return "aktivitas_gambar"
}

// Model kategori kegiatan
type KategoriKegiatan struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Nama string `json:"nama"`
}

func (KategoriKegiatan) TableName() string {
	return "kategori_kegiatan"
}
