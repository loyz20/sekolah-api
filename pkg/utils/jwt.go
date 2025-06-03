package utils

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(GetEnv("JWT_SECRET", "default-secret")) // Ganti dgn utils.GetEnv jika mau dari ENV

type JWTClaims struct {
	PenggunaID string `json:"pengguna_id"`
	SekolahID  string `json:"sekolah_id"`
	PeranIDStr string `json:"peran_id_str"`
	jwt.RegisteredClaims
}

func GenerateToken(penggunaID, sekolahID, peranID string) (string, error) {
	claims := JWTClaims{
		PenggunaID: penggunaID,
		SekolahID:  sekolahID,
		PeranIDStr: peranID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func GetSecret() []byte {
	return jwtSecret
}

func GenerateRefreshToken() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
