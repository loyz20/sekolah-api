package repository

import (
	"sekolah-api/internal/auth/domain"
)

type RefreshTokenRepository interface {
	Store(token *domain.RefreshToken) error
	FindByToken(token string) (*domain.RefreshToken, error)
	Delete(token string) error
	DeleteByPenggunaID(penggunaID string) error
}
