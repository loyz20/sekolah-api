package persistence

import (
	"context"
	"sekolah-api/internal/pengguna/domain"
	"sekolah-api/internal/pengguna/repository"

	"gorm.io/gorm"
)

type penggunaRepo struct {
	db *gorm.DB
}

func NewPenggunaRepository(db *gorm.DB) repository.PenggunaRepository {
	return &penggunaRepo{db: db}
}

func (r *penggunaRepo) FindByUsername(ctx context.Context, username string) (*domain.Pengguna, error) {
	var user domain.Pengguna
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *penggunaRepo) Save(ctx context.Context, user *domain.Pengguna) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *penggunaRepo) FindByID(id string) (*domain.Pengguna, error) {
	var pengguna domain.Pengguna
	if err := r.db.Where("pengguna_id = ?", id).First(&pengguna).Error; err != nil {
		return nil, err
	}
	return &pengguna, nil
}

func (r *penggunaRepo) FindAll(ctx context.Context, limit, offset int) ([]domain.Pengguna, int, error) {
	var users []domain.Pengguna
	var count int64

	if err := r.db.WithContext(ctx).Model(&domain.Pengguna{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Order("created_at desc").
		Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, int(count), nil
}

func (r *penggunaRepo) Update(ctx context.Context, user *domain.Pengguna) error {
	return r.db.WithContext(ctx).
		Model(&domain.Pengguna{}).
		Where("pengguna_id = ?", user.PenggunaID).
		Updates(user).Error
}

func (r *penggunaRepo) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("pengguna_id = ?", id).Delete(&domain.Pengguna{}).Error
}
