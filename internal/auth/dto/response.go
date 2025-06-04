// --- internal/pengguna/dto/response.go ---
package dto

import "time"

type UserProfile struct {
	PenggunaID     string  `json:"pengguna_id"`
	Nama           string  `json:"nama"`
	PeranIDStr     string  `json:"peran_id_str"`
	SekolahID      string  `json:"sekolah_id"`
	PTKID          *string `json:"ptk_id"`
	PesertaDidikID *string `json:"peserta_didik_id"`
}

type LoginResponse struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	ExpiredAt    time.Time   `json:"expired_at"`
	User         UserProfile `json:"user"`
}
