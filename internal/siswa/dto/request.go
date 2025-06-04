package dto

import "time"

type CreateSiswaRequest struct {
	NIPD                string    `json:"nipd" binding:"required"`
	TanggalMasukSekolah time.Time `json:"tanggal_masuk_sekolah" binding:"required"`
	SekolahAsal         string    `json:"sekolah_asal"`

	PesertaDidikID string    `json:"peserta_didik_id" binding:"required,uuid"`
	Nama           string    `json:"nama" binding:"required"`
	NISN           string    `json:"nisn"`
	JenisKelamin   string    `json:"jenis_kelamin" binding:"required,oneof=L P"`
	NIK            string    `json:"nik"`
	TempatLahir    string    `json:"tempat_lahir"`
	TanggalLahir   time.Time `json:"tanggal_lahir"`

	AgamaIDStr string `json:"agama_id_str"`

	SemesterID          string `json:"semester_id"`
	AnggotaRombelID     string `json:"anggota_rombel_id"`
	RombonganBelajarID  string `json:"rombongan_belajar_id"`
	TingkatPendidikanID string `json:"tingkat_pendidikan_id"`
	KurikulumID         int    `json:"kurikulum_id"`
	KebutuhanKhusus     string `json:"kebutuhan_khusus"`
}

type UpdateSiswaRequest struct {
	NIPD                string    `json:"nipd"`
	TanggalMasukSekolah time.Time `json:"tanggal_masuk_sekolah"`
	SekolahAsal         string    `json:"sekolah_asal"`
	Nama                string    `json:"nama"`
	NISN                string    `json:"nisn"`
	JenisKelamin        string    `json:"jenis_kelamin"`
	NIK                 string    `json:"nik"`
	TempatLahir         string    `json:"tempat_lahir"`
	TanggalLahir        time.Time `json:"tanggal_lahir"`
	AgamaIDStr          string    `json:"agama_id_str"`
	NomorTeleponRumah   string    `json:"nomor_telepon_rumah"`
	NomorTeleponSeluler string    `json:"nomor_telepon_seluler"`
	NamaAyah            string    `json:"nama_ayah"`
	PekerjaanAyahID     int       `json:"pekerjaan_ayah_id"`
	NamaIbu             string    `json:"nama_ibu"`
	PekerjaanIbuID      int       `json:"pekerjaan_ibu_id"`
	NamaWali            string    `json:"nama_wali"`
	PekerjaanWaliID     int       `json:"pekerjaan_wali_id"`
	AnakKeberapa        string    `json:"anak_keberapa"`
	TinggiBadan         string    `json:"tinggi_badan"`
	BeratBadan          string    `json:"berat_badan"`
	Email               string    `json:"email"`
	SemesterID          string    `json:"semester_id"`
	AnggotaRombelID     string    `json:"anggota_rombel_id"`
	RombonganBelajarID  string    `json:"rombongan_belajar_id"`
	TingkatPendidikanID string    `json:"tingkat_pendidikan_id"`
	KurikulumID         int       `json:"kurikulum_id"`
	KebutuhanKhusus     string    `json:"kebutuhan_khusus"`
}
