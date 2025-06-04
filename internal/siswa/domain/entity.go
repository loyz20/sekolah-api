package domain

import (
	"time"
)

type Siswa struct {
	RegistrasiID          string    `gorm:"primaryKey" json:"registrasi_id"`
	JenisPendaftaranID    string    `gorm:"type:varchar(5)" json:"jenis_pendaftaran_id"`
	JenisPendaftaranIDStr string    `gorm:"-" json:"jenis_pendaftaran_id_str"` // view only
	NIPD                  string    `gorm:"type:varchar(20)" json:"nipd"`
	TanggalMasukSekolah   time.Time `json:"tanggal_masuk_sekolah"`
	SekolahAsal           string    `gorm:"type:varchar(100)" json:"sekolah_asal"`

	PesertaDidikID string    `gorm:"unique" json:"peserta_didik_id"`
	Nama           string    `gorm:"type:varchar(100)" json:"nama"`
	NISN           string    `gorm:"type:varchar(20)" json:"nisn"`
	JenisKelamin   string    `gorm:"type:char(1)" json:"jenis_kelamin"`
	NIK            string    `gorm:"type:varchar(20)" json:"nik"`
	TempatLahir    string    `gorm:"type:varchar(50)" json:"tempat_lahir"`
	TanggalLahir   time.Time `json:"tanggal_lahir"`

	AgamaID    int    `json:"agama_id"`
	AgamaIDStr string `gorm:"column:agamar_id_str" json:"agama_id_str"`

	NomorTeleponRumah   string `gorm:"type:varchar(20)" json:"nomor_telepon_rumah"`
	NomorTeleponSeluler string `gorm:"type:varchar(20)" json:"nomor_telepon_seluler"`

	NamaAyah           string `gorm:"type:varchar(100)" json:"nama_ayah"`
	PekerjaanAyahID    int    `json:"pekerjaan_ayah_id"`
	PekerjaanAyahIDStr string `gorm:"-" json:"pekerjaan_ayah_id_str"`

	NamaIbu           string `gorm:"type:varchar(100)" json:"nama_ibu"`
	PekerjaanIbuID    int    `json:"pekerjaan_ibu_id"`
	PekerjaanIbuIDStr string `gorm:"-" json:"pekerjaan_ibu_id_str"`

	NamaWali           string `gorm:"type:varchar(100)" json:"nama_wali"`
	PekerjaanWaliID    int    `json:"pekerjaan_wali_id"`
	PekerjaanWaliIDStr string `gorm:"-" json:"pekerjaan_wali_id_str"`

	AnakKeberapa string `gorm:"type:varchar(5)" json:"anak_keberapa"`
	TinggiBadan  string `gorm:"type:varchar(5)" json:"tinggi_badan"`
	BeratBadan   string `gorm:"type:varchar(5)" json:"berat_badan"`

	Email string `gorm:"type:varchar(100)" json:"email"`

	SemesterID         string `gorm:"type:varchar(10)" json:"semester_id"`
	AnggotaRombelID    string `gorm:"type:uuid" json:"anggota_rombel_id"`
	RombonganBelajarID string `gorm:"type:uuid" json:"rombongan_belajar_id"`
	NamaRombel         string `gorm:"-" json:"nama_rombel"` // view only

	TingkatPendidikanID string `gorm:"type:varchar(10)" json:"tingkat_pendidikan_id"`

	KurikulumID    int    `json:"kurikulum_id"`
	KurikulumIDStr string `json:"kurikulum_id_str"`

	KebutuhanKhusus string `gorm:"type:varchar(100)" json:"kebutuhan_khusus"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
