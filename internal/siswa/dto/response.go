package dto

import "time"

type SiswaResponse struct {
	RegistrasiID          string    `json:"registrasi_id"`
	JenisPendaftaranID    string    `json:"jenis_pendaftaran_id"`
	JenisPendaftaranIDStr string    `json:"jenis_pendaftaran_id_str"`
	NIPD                  string    `json:"nipd"`
	TanggalMasukSekolah   time.Time `json:"tanggal_masuk_sekolah"`
	SekolahAsal           string    `json:"sekolah_asal"`
	PesertaDidikID        string    `json:"peserta_didik_id"`
	Nama                  string    `json:"nama"`
	NISN                  string    `json:"nisn"`
	JenisKelamin          string    `json:"jenis_kelamin"`
	NIK                   string    `json:"nik"`
	TempatLahir           string    `json:"tempat_lahir"`
	TanggalLahir          time.Time `json:"tanggal_lahir"`
	AgamaID               int       `json:"agama_id"`
	AgamaIDStr            string    `json:"agama_id_str"`
	NomorTeleponRumah     string    `json:"nomor_telepon_rumah"`
	NomorTeleponSeluler   string    `json:"nomor_telepon_seluler"`
	NamaAyah              string    `json:"nama_ayah"`
	PekerjaanAyahID       int       `json:"pekerjaan_ayah_id"`
	PekerjaanAyahIDStr    string    `json:"pekerjaan_ayah_id_str"`
	NamaIbu               string    `json:"nama_ibu"`
	PekerjaanIbuID        int       `json:"pekerjaan_ibu_id"`
	PekerjaanIbuIDStr     string    `json:"pekerjaan_ibu_id_str"`
	NamaWali              string    `json:"nama_wali"`
	PekerjaanWaliID       int       `json:"pekerjaan_wali_id"`
	PekerjaanWaliIDStr    string    `json:"pekerjaan_wali_id_str"`
	AnakKeberapa          string    `json:"anak_keberapa"`
	TinggiBadan           string    `json:"tinggi_badan"`
	BeratBadan            string    `json:"berat_badan"`
	Email                 string    `json:"email"`
	SemesterID            string    `json:"semester_id"`
	AnggotaRombelID       string    `json:"anggota_rombel_id"`
	RombonganBelajarID    string    `json:"rombongan_belajar_id"`
	NamaRombel            string    `json:"nama_rombel"`
	TingkatPendidikanID   string    `json:"tingkat_pendidikan_id"`
	KurikulumID           int       `json:"kurikulum_id"`
	KurikulumIDStr        string    `json:"kurikulum_id_str"`
	KebutuhanKhusus       string    `json:"kebutuhan_khusus"`
}
