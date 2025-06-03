package usecase

import (
	"context"
	"sekolah-api/internal/pengguna/domain"
	"sekolah-api/internal/pengguna/dto"
	"sekolah-api/internal/pengguna/repository"
	"sekolah-api/pkg/utils"

	"github.com/google/uuid"
)

type PenggunaUsecase struct {
	repo repository.PenggunaRepository
}

func NewPenggunaUsecase(r repository.PenggunaRepository) *PenggunaUsecase {
	return &PenggunaUsecase{repo: r}
}

func (u *PenggunaUsecase) Create(ctx context.Context, req dto.CreatePenggunaRequest) error {
	hashed, _ := utils.HashPassword(req.Password)
	pengguna := domain.Pengguna{
		PenggunaID:     uuid.NewString(),
		SekolahID:      req.SekolahID,
		Username:       req.Username,
		Nama:           req.Nama,
		PeranIDStr:     req.PeranIDStr,
		Password:       hashed,
		Alamat:         req.Alamat,
		NoTelepon:      req.NoTelepon,
		NoHP:           req.NoHP,
		PtkID:          req.PtkID,
		PesertaDidikID: req.PesertaDidikID,
	}
	return u.repo.Save(ctx, &pengguna)
}

func (u *PenggunaUsecase) Update(ctx context.Context, id string, req dto.UpdatePenggunaRequest) error {
	pengguna, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}

	pengguna.Nama = req.Nama
	pengguna.Alamat = req.Alamat
	pengguna.NoTelepon = req.NoTelepon
	pengguna.NoHP = req.NoHP

	return u.repo.Update(ctx, pengguna)
}

func (u *PenggunaUsecase) GetAll(ctx context.Context, page, limit int) ([]domain.Pengguna, int, error) {
	offset := (page - 1) * limit
	return u.repo.FindAll(ctx, limit, offset)
}

func (u *PenggunaUsecase) GetByID(ctx context.Context, id string) (*domain.Pengguna, error) {
	return u.repo.FindByID(id)
}

func (u *PenggunaUsecase) Delete(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}
