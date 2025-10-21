// package models

// import "time"

// // ==============================
// // 🧩 Model: Mitra
// // ==============================
// type Mitra struct {
// 	ID              uint           `json:"id"`
// 	Nama            string         `json:"nama"`
// 	Logo            *string        `json:"logo,omitempty"`
// 	Deskripsi       *string        `json:"deskripsi,omitempty"`
// 	TanggalMulai    *time.Time     `json:"tanggal_mulai,omitempty"`
// 	TanggalSelesai  *time.Time     `json:"tanggal_selesai,omitempty"`
// 	KategoriMitraID *uint          `json:"kategori_mitra_id"`
// 	KategoriMitra   *KategoriMitra `json:"kategori_mitra,omitempty" gorm:"foreignKey:KategoriMitraID"`
// 	CreatedAt       time.Time      `json:"created_at"`
// 	UpdatedAt       time.Time      `json:"updated_at"`
// }

// func (Mitra) TableName() string {
// 	return "mitra"
// }

// // ==============================
// // 🧩 Model: Kategori Mitra
// // ==============================
// type KategoriMitra struct {
// 	ID         uint   `json:"id"`
// 	Nama       string `json:"nama"`
// 	Deskripsi  string `json:"deskripsi"`
// 	Keterangan string `json:"keterangan"`
// }

// func (KategoriMitra) TableName() string {
// 	return "kategori_mitra"
// }

// // ==============================
// // 💰 Model: Donasi
// // ==============================
// // type Donasi struct {
// // 	ID            uint      `json:"id" gorm:"primaryKey"`
// // 	NamaDonatur   string    `json:"nama_donatur"`
// // 	JumlahDonasi  float64   `json:"jumlah_donasi"`
// // 	Pesan         *string   `json:"pesan,omitempty"`
// // 	BuktiTransfer *string   `json:"bukti_transfer,omitempty"`
// // 	Status        string    `json:"status" gorm:"default:'pending'"` // pending, verified, rejected
// // 	MitraID       *uint     `json:"mitra_id,omitempty"`
// // 	Mitra         *Mitra    `json:"mitra,omitempty" gorm:"foreignKey:MitraID"`
// // 	CreatedAt     time.Time `json:"created_at"`
// // 	UpdatedAt     time.Time `json:"updated_at"`
// // }

// type Donasi struct {
// 	ID      uint   `json:"id" gorm:"primaryKey"`
// 	MitraID uint   `json:"mitra_id"`
// 	Jumlah  string `json:"jumlah"`
// 	Pesan   string `json:"pesan"`
// 	Metode  string `json:"metode"`
// 	Status  string `json:"status"`
// 	Mitra   Mitra  `json:"mitra" gorm:"foreignKey:MitraID"`
// }

// func (Donasi) TableName() string {
// 	return "donasi"
// }

package models

import "time"

// ==============================
// 🧩 Model: Mitra
// ==============================
type Mitra struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Nama            string         `json:"nama"`
	Logo            *string        `json:"logo,omitempty"`
	Deskripsi       *string        `json:"deskripsi,omitempty"`
	TanggalMulai    *time.Time     `json:"tanggal_mulai,omitempty"`
	TanggalSelesai  *time.Time     `json:"tanggal_selesai,omitempty"`
	KategoriMitraID *uint          `json:"kategori_mitra_id"`
	KategoriMitra   *KategoriMitra `json:"kategori_mitra,omitempty" gorm:"foreignKey:KategoriMitraID;references:ID"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`

	// Relasi ke Donasi
	Donasis []Donasi `json:"donasis" gorm:"foreignKey:MitraID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Mitra) TableName() string {
	return "mitra"
}

// ==============================
// 🧩 Model: Kategori Mitra
// ==============================
type KategoriMitra struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Nama       string `json:"nama"`
	Deskripsi  string `json:"deskripsi"`
	Keterangan string `json:"keterangan"`
}

func (KategoriMitra) TableName() string {
	return "kategori_mitra"
}

// ==============================
// 💰 Model: Donasi
// ==============================
// type Donasi struct {
// 	ID        uint      `json:"id" gorm:"primaryKey"`
// 	MitraID   uint      `json:"mitra_id"`
// 	Nama      string    `json:"nama"` // 🔹 Tambahkan kolom nama
// 	Jumlah    string    `json:"jumlah"`
// 	Pesan     string    `json:"pesan"`
// 	Metode    string    `json:"metode"`
// 	Status    string    `json:"status" gorm:"default:'pending'"`
// 	Mitra     *Mitra    `json:"mitra,omitempty" gorm:"foreignKey:MitraID;references:ID"`
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// }

//	func (Donasi) TableName() string {
//		return "donasi"
//	}
type Donasi struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	MitraID       uint      `json:"mitra_id"`
	Nama          string    `json:"nama"`
	Jumlah        string    `json:"jumlah"`
	Pesan         string    `json:"pesan"`
	Metode        string    `json:"metode"`
	Status        string    `json:"status" gorm:"default:'pending'"`
	BuktiTransfer *string   `json:"bukti_transfer,omitempty"`
	Mitra         *Mitra    `json:"mitra,omitempty" gorm:"foreignKey:MitraID;references:ID"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (Donasi) TableName() string {
	return "donasi"
}
