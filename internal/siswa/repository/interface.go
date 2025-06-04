package repository

import (
	"context"
	"sekolah-api/internal/siswa/domain"
)

type SiswaRepository interface {
	FindByID(id string) (*domain.Siswa, error)
	FindByUsername(ctx context.Context, username string) (*domain.Siswa, error)
	FindAll(ctx context.Context, limit, offset int) ([]domain.Siswa, int, error)
	Save(ctx context.Context, user *domain.Siswa) error
	Update(ctx context.Context, user *domain.Siswa) error
	Delete(ctx context.Context, id string) error
}
