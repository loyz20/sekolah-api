package repository

import (
	"context"
	"sekolah-api/internal/pengguna/domain"
)

type PenggunaRepository interface {
	FindByID(id string) (*domain.Pengguna, error)
	FindByUsername(ctx context.Context, username string) (*domain.Pengguna, error)
	FindAll(ctx context.Context, limit, offset int) ([]domain.Pengguna, int, error)
	Save(ctx context.Context, user *domain.Pengguna) error
	Update(ctx context.Context, user *domain.Pengguna) error
	Delete(ctx context.Context, id string) error
}
