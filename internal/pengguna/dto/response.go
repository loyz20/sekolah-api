package dto

type PenggunaResponse struct {
    PenggunaID      string  `json:"pengguna_id"`
    SekolahID       string  `json:"sekolah_id"`
    Username        string  `json:"username"`
    Nama            string  `json:"nama"`
    PeranIDStr      string  `json:"peran_id_str"`
    Alamat          string  `json:"alamat"`
    NoTelepon       string  `json:"no_telepon"`
    NoHP            string  `json:"no_hp"`
    PtkID           *string `json:"ptk_id"`
    PesertaDidikID  *string `json:"peserta_didik_id"`
}
