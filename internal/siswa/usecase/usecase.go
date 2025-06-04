package usecase

import (
	"context"
	"sekolah-api/internal/siswa/domain"
	"sekolah-api/internal/siswa/dto"
	"sekolah-api/internal/siswa/repository"

	"github.com/google/uuid"
)

type SiswaUsecase struct {
	repo repository.SiswaRepository
}

func NewSiswaUsecase(r repository.SiswaRepository) *SiswaUsecase {
	return &SiswaUsecase{repo: r}
}

func (u *SiswaUsecase) Create(ctx context.Context, req dto.CreateSiswaRequest) error {
	siswa := domain.Siswa{
		RegistrasiID:        uuid.NewString(),
		NIPD:                req.NIPD,
		TanggalMasukSekolah: req.TanggalMasukSekolah,
		SekolahAsal:         req.SekolahAsal,
		PesertaDidikID:      req.PesertaDidikID,
		Nama:                req.Nama,
		NISN:                req.NISN,
		JenisKelamin:        req.JenisKelamin,
		NIK:                 req.NIK,
		TempatLahir:         req.TempatLahir,
		TanggalLahir:        req.TanggalLahir,
		AgamaIDStr:          req.AgamaIDStr,
	}
	return u.repo.Save(ctx, &siswa)
}

func (u *SiswaUsecase) Update(ctx context.Context, id string, req dto.UpdateSiswaRequest) error {
	siswa, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}
	siswa.NIPD = req.NIPD
	siswa.TanggalMasukSekolah = req.TanggalMasukSekolah
	siswa.Nama = req.Nama
	siswa.NISN = req.NISN
	siswa.JenisKelamin = req.JenisKelamin
	siswa.TempatLahir = req.TempatLahir
	siswa.TanggalLahir = req.TanggalLahir
	siswa.AgamaIDStr = req.AgamaIDStr

	return u.repo.Update(ctx, siswa)
}

func (u *SiswaUsecase) GetAll(ctx context.Context, page, limit int) ([]domain.Siswa, int, error) {
	offset := (page - 1) * limit
	return u.repo.FindAll(ctx, limit, offset)
}

func (u *SiswaUsecase) GetByID(ctx context.Context, id string) (*domain.Siswa, error) {
	return u.repo.FindByID(id)
}

func (u *SiswaUsecase) Delete(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}
