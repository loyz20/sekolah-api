package persistence

import (
	"context"
	"sekolah-api/internal/siswa/domain"
	"sekolah-api/internal/siswa/repository"

	"gorm.io/gorm"
)

type siswaRepo struct {
	db *gorm.DB
}

func NewSiswaRepository(db *gorm.DB) repository.SiswaRepository {
	return &siswaRepo{db: db}
}

func (r *siswaRepo) FindByUsername(ctx context.Context, username string) (*domain.Siswa, error) {
	var user domain.Siswa
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *siswaRepo) Save(ctx context.Context, user *domain.Siswa) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *siswaRepo) FindByID(id string) (*domain.Siswa, error) {
	var siswa domain.Siswa
	if err := r.db.Where("siswa_id = ?", id).First(&siswa).Error; err != nil {
		return nil, err
	}
	return &siswa, nil
}

func (r *siswaRepo) FindAll(ctx context.Context, limit, offset int) ([]domain.Siswa, int, error) {
	var users []domain.Siswa
	var count int64

	if err := r.db.WithContext(ctx).Model(&domain.Siswa{}).Count(&count).Error; err != nil {
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

func (r *siswaRepo) Update(ctx context.Context, user *domain.Siswa) error {
	return r.db.WithContext(ctx).
		Model(&domain.Siswa{}).
		Where("siswa_id = ?", user.PesertaDidikID).
		Updates(user).Error
}

func (r *siswaRepo) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("siswa_id = ?", id).Delete(&domain.Siswa{}).Error
}
