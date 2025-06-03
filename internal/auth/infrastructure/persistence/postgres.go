// --- internal/pengguna/infrastructure/persistence/postgres.go ---
package persistence

import (
	"sekolah-api/internal/auth/domain"
	"sekolah-api/internal/auth/repository"

	"gorm.io/gorm"
)

type refreshTokenRepo struct {
	db *gorm.DB
}

func NewRefreshTokenRepo(db *gorm.DB) repository.RefreshTokenRepository {
	return &refreshTokenRepo{db}
}

func (r *refreshTokenRepo) Store(token *domain.RefreshToken) error {
	return r.db.Create(token).Error
}

func (r *refreshTokenRepo) FindByToken(token string) (*domain.RefreshToken, error) {
	var rt domain.RefreshToken
	err := r.db.Where("token = ?", token).First(&rt).Error
	if err != nil {
		return nil, err
	}
	return &rt, nil
}

func (r *refreshTokenRepo) Delete(token string) error {
	return r.db.Where("token = ?", token).Delete(&domain.RefreshToken{}).Error
}

func (r *refreshTokenRepo) DeleteByPenggunaID(penggunaID string) error {
	return r.db.Where("pengguna_id = ?", penggunaID).Delete(&domain.RefreshToken{}).Error
}
