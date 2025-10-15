package models

import "time"

type StrukturOrganisasi struct {
	ID        uint                `gorm:"primaryKey" json:"id"`
	Nama      string              `json:"nama"`
	JabatanID *uint               `json:"jabatan_id"`
	Jabatan   *Jabatan            `json:"jabatan,omitempty" gorm:"foreignKey:JabatanID"`
	ParentID  *uint               `json:"parent_id,omitempty"`
	Parent    *StrukturOrganisasi `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Foto      *string             `json:"foto,omitempty"`
	Deskripsi *string             `json:"deskripsi,omitempty"`
	Urutan    int                 `json:"urutan"`
	Status    bool                `json:"status"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
}

func (StrukturOrganisasi) TableName() string {
	return "struktur_organisasis"
}

type Jabatan struct {
	ID   uint   `json:"id"`
	Nama string `json:"nama"`
}

func (Jabatan) TableName() string {
	return "jabatan"
}
