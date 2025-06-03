package seeder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sekolah-api/internal/pengguna/domain"
	"sekolah-api/pkg/utils"

	"gorm.io/gorm"
)

func SeedPengguna(db *gorm.DB) {
	file, err := os.Open("pkg/data/pengguna.json")
	if err != nil {
		log.Fatalf("Gagal membuka file JSON: %v", err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var penggunaList []domain.Pengguna
	if err := json.Unmarshal(byteValue, &penggunaList); err != nil {
		log.Fatalf("Gagal parse JSON: %v", err)
	}

	for _, pengguna := range penggunaList {
		// Cek apakah sudah ada
		var existing domain.Pengguna
		err := db.Where("username = ?", pengguna.Username).First(&existing).Error
		if err == gorm.ErrRecordNotFound {
			if pengguna.PeranIDStr == "PTK" {
				// Hash password plaintext dari JSON
				hashed, err := utils.HashPassword("123456")
				if err != nil {
					log.Printf("Gagal hash password untuk %s: %v", pengguna.Username, err)
					continue
				}
				pengguna.Password = hashed

				if err := db.Create(&pengguna).Error; err != nil {
					log.Printf("Gagal menyimpan pengguna %s: %v", pengguna.Username, err)
				} else {
					log.Printf("✅ Pengguna %s disimpan", pengguna.Username)
				}
			} else {
				continue
			}

		} else {
			log.Printf("ℹ️  Pengguna %s sudah ada", pengguna.Username)
		}
	}
}
