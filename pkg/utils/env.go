package utils

import "os"

// GetEnv membaca variabel env dan fallback jika tidak ditemukan
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
