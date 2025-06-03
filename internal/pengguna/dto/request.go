package dto

type CreatePenggunaRequest struct {
    SekolahID      string  `json:"sekolah_id" binding:"required,uuid"`
    Username       string  `json:"username" binding:"required,email"`
    Nama           string  `json:"nama" binding:"required"`
    PeranIDStr     string  `json:"peran_id_str" binding:"required"`
    Password       string  `json:"password" binding:"required,min=6"`
    Alamat         string  `json:"alamat"`
    NoTelepon      string  `json:"no_telepon"`
    NoHP           string  `json:"no_hp"`
    PtkID          *string `json:"ptk_id"`
    PesertaDidikID *string `json:"peserta_didik_id"`
}

type UpdatePenggunaRequest struct {
    Nama           string  `json:"nama"`
    Alamat         string  `json:"alamat"`
    NoTelepon      string  `json:"no_telepon"`
    NoHP           string  `json:"no_hp"`
}
