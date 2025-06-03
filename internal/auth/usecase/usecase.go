package usecase

import (
	"errors"
	"fmt"
	"sekolah-api/internal/auth/domain"
	"sekolah-api/internal/auth/dto"
	authRepo "sekolah-api/internal/auth/repository"
	pengRepo "sekolah-api/internal/pengguna/repository"
	"sekolah-api/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthUsecase struct {
	PenggunaRepo     pengRepo.PenggunaRepository
	RefreshTokenRepo authRepo.RefreshTokenRepository
}

func NewAuthUsecase(pRepo pengRepo.PenggunaRepository, rRepo authRepo.RefreshTokenRepository) *AuthUsecase {
	return &AuthUsecase{
		PenggunaRepo:     pRepo,
		RefreshTokenRepo: rRepo,
	}
}

// Login logic
func (u *AuthUsecase) Login(ctx *gin.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	pengguna, err := u.PenggunaRepo.FindByUsername(ctx, req.Username)
	if err != nil || !utils.CheckPasswordHash(req.Password, pengguna.Password) {
		return nil, errors.New("username atau password salah")
	}

	accessToken, err := utils.GenerateToken(pengguna.PenggunaID, pengguna.SekolahID, pengguna.PeranIDStr)
	if err != nil {
		return nil, err
	}

	refreshToken := uuid.NewString()
	expiresAt := time.Now().Add(7 * 24 * time.Hour)

	refresh := &domain.RefreshToken{
		ID:         uuid.NewString(),
		PenggunaID: pengguna.PenggunaID,
		Token:      refreshToken,
		ExpiresAt:  expiresAt,
		CreatedAt:  time.Now(),
	}

	if err := u.RefreshTokenRepo.Store(refresh); err != nil {
		return nil, fmt.Errorf("gagal simpan refresh token: %w", err)
	}

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiredAt:    expiresAt,
		User: dto.UserProfile{
			PenggunaID: pengguna.PenggunaID,
			Nama:       pengguna.Nama,
			PeranIDStr: pengguna.PeranIDStr,
			SekolahID:  pengguna.SekolahID,
		},
	}, nil
}

// Refresh Token
func (u *AuthUsecase) RefreshToken(tokenStr string) (string, error) {
	refresh, err := u.RefreshTokenRepo.FindByToken(tokenStr)
	if err != nil || refresh.ExpiresAt.Before(time.Now()) {
		return "", errors.New("refresh token tidak valid atau kadaluarsa")
	}

	pengguna, err := u.PenggunaRepo.FindByID(refresh.PenggunaID)
	if err != nil {
		return "", errors.New("pengguna tidak ditemukan")
	}

	accessToken, err := utils.GenerateToken(pengguna.PenggunaID, pengguna.SekolahID, pengguna.PeranIDStr)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

// Logout
func (u *AuthUsecase) Logout(token string) error {
	return u.RefreshTokenRepo.Delete(token)
}

// Ganti Password
func (u *AuthUsecase) ChangePassword(ctx *gin.Context, penggunaID, oldPwd, newPwd string) error {
	user, err := u.PenggunaRepo.FindByID(penggunaID)
	if err != nil {
		return errors.New("pengguna tidak ditemukan")
	}

	if !utils.CheckPasswordHash(oldPwd, user.Password) {
		return errors.New("password lama salah")
	}

	hashed, _ := utils.HashPassword(newPwd)
	user.Password = hashed

	if err := u.PenggunaRepo.Update(ctx, user); err != nil {
		return errors.New("gagal mengubah password")
	}

	_ = u.RefreshTokenRepo.DeleteByPenggunaID(penggunaID)
	return nil
}
