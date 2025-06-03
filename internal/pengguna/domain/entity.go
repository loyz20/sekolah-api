package domain

import "time"

type Pengguna struct {
	PenggunaID     string  `gorm:"type:uuid;primary_key" json:"pengguna_id"`
	SekolahID      string  `gorm:"type:uuid" json:"sekolah_id"`
	Username       string  `json:"username"`
	Nama           string  `json:"nama"`
	PeranIDStr     string  `json:"peran_id_str"`
	Password       string  `json:"-"`
	Alamat         string  `json:"alamat"`
	NoTelepon      string  `json:"no_telepon"`
	NoHP           string  `json:"no_hp"`
	PtkID          *string `json:"ptk_id"`
	PesertaDidikID *string `json:"peserta_didik_id"`
	CreatedAt      time.Time
	UpdateAt       time.Time
}
