package seeder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	penggunaDomain "sekolah-api/internal/pengguna/domain"
	"sekolah-api/internal/siswa/domain"
	"sekolah-api/pkg/utils"
	"time"

	"gorm.io/gorm"
)

const layout = "2006-01-02"

type rawSiswa struct {
	RegistrasiID        string `json:"registrasi_id"`
	PesertaDidikID      string `json:"peserta_didik_id"`
	Nama                string `json:"nama"`
	TanggalLahir        string `json:"tanggal_lahir"`
	TempatLahir         string `json:"tempat_lahir"`
	TanggalMasukSekolah string `json:"tanggal_masuk_sekolah"`

	NIPD                string `json:"nipd"`
	JenisKelamin        string `json:"jenis_kelamin"`
	Nisn                string `json:"nisn"`
	Nik                 string `json:"nik"`
	AgamaIDStr          string `json:"agama_id_str"`
	AnggotaRombelID     string `json:"anggota_rombel_id"`
	RombonganBelajarID  string `json:"rombongan_belajar_id"`
	TingkatPendidikanID string `json:"tingkat_pendidikan_id"`
	KebutuhanKhusus     string `json:"kebutuhan_khusus"`
}

func SeedSiswa(db *gorm.DB) {
	file, err := os.Open("pkg/data/siswa.json")
	if err != nil {
		log.Fatalf("❌ Gagal membuka file JSON: %v", err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var siswaList []rawSiswa
	if err := json.Unmarshal(byteValue, &siswaList); err != nil {
		log.Fatalf("❌ Gagal parse JSON: %v", err)
	}

	for _, raw := range siswaList {
		tglLahir, err := time.Parse(layout, raw.TanggalLahir)
		if err != nil {
			log.Printf("❌ Format tanggal_lahir tidak valid untuk %s: %v", raw.Nama, err)
			continue
		}

		tglMasuk, err := time.Parse(layout, raw.TanggalMasukSekolah)
		if err != nil {
			log.Printf("❌ Format tanggal_masuk_sekolah tidak valid untuk %s: %v", raw.Nama, err)
			continue
		}

		siswa := domain.Siswa{
			RegistrasiID:        raw.RegistrasiID,
			PesertaDidikID:      raw.PesertaDidikID,
			Nama:                raw.Nama,
			TanggalLahir:        tglLahir,
			TempatLahir:         raw.TempatLahir,
			TanggalMasukSekolah: tglMasuk,
			NIPD:                raw.NIPD,
			NIK:                 raw.Nik,
			JenisKelamin:        raw.JenisKelamin,
			NISN:                raw.Nisn,
			AgamaIDStr:          raw.AgamaIDStr,
			AnggotaRombelID:     raw.AnggotaRombelID,
			RombonganBelajarID:  raw.RombonganBelajarID,
			TingkatPendidikanID: raw.TingkatPendidikanID,
			KebutuhanKhusus:     raw.KebutuhanKhusus,
		}

		var existing domain.Siswa
		err = db.Where("peserta_didik_id = ?", siswa.PesertaDidikID).First(&existing).Error
		if err == gorm.ErrRecordNotFound {
			// ✅ Simpan siswa
			if err := db.Create(&siswa).Error; err != nil {
				log.Printf("❌ Gagal insert siswa %s: %v", siswa.Nama, err)
				continue
			}

			// ✅ Tambahkan pengguna untuk siswa
			hashedPassword, err := utils.HashPassword("123456")
			if err != nil {
				log.Printf("❌ Gagal hash password untuk %s: %v", siswa.Nama, err)
				continue
			}

			pengguna := penggunaDomain.Pengguna{
				PenggunaID:     siswa.PesertaDidikID,
				SekolahID:      "b9dd9b09-dda7-4251-9cec-c2f06db7bbf9",
				Nama:           siswa.Nama,
				Username:       siswa.NIPD,
				Password:       hashedPassword,
				PeranIDStr:     "siswa",
				PesertaDidikID: ptr(siswa.PesertaDidikID),
			}

			if err := db.Create(&pengguna).Error; err != nil {
				log.Printf("❌ Gagal membuat akun pengguna untuk %s: %v", siswa.Nama, err)
			} else {
				log.Printf("✅ Berhasil buat akun pengguna untuk siswa: %s", siswa.Nama)
			}
		} else if err != nil {
			log.Printf("❌ Gagal cek siswa %s: %v", siswa.Nama, err)
		} else {
			log.Printf("ℹ️  Siswa sudah ada: %s", siswa.Nama)
		}
	}
}

func ptr[T any](v T) *T {
	return &v
}
