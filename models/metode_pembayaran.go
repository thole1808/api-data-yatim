// models/metode_pembayaran.go
package models

type MetodePembayaran struct {
	ID                uint             `json:"id" gorm:"primaryKey"`
	Nama              string           `json:"nama"`
	IDJenisPembayaran uint             `json:"id_jenis_pembayaran"`
	NomorRekening     string           `json:"nomor_rekening"`
	NamaPemilik       string           `json:"nama_pemilik"`
	QRImage           string           `json:"qr_image"`
	Deskripsi         string           `json:"deskripsi"`
	Status            string           `json:"status"`
	JenisPembayaran   *JenisPembayaran `json:"jenis_pembayaran" gorm:"foreignKey:IDJenisPembayaran"`
}

func (MetodePembayaran) TableName() string {
	return "metode_pembayaran"
}
