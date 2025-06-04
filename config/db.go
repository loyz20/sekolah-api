package config

import (
	"fmt"
	"log"
	authDom "sekolah-api/internal/auth/domain"
	pengDom "sekolah-api/internal/pengguna/domain"
	siswaDom "sekolah-api/internal/siswa/domain"
	"sekolah-api/pkg/utils"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := utils.GetEnv("DB_HOST", "localhost")
	port := utils.GetEnv("DB_PORT", "5432")
	user := utils.GetEnv("DB_USER", "postgres")
	password := utils.GetEnv("DB_PASSWORD", "postgres")
	dbname := utils.GetEnv("DB_NAME", "sekolah_db")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatalf("❌ gagal koneksi database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ gagal dapatkan instance DB: %v", err)
	}

	// Optional: set connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("✅ berhasil koneksi ke database")
	db.AutoMigrate(
		&pengDom.Pengguna{},
		&authDom.RefreshToken{},
		&siswaDom.Siswa{},
	)

	DB = db
	return db
}
