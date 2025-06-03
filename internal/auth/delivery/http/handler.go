package http

import (
	"net/http"
	"sekolah-api/internal/auth/dto"
	"sekolah-api/internal/auth/usecase"
	"sekolah-api/pkg/response"
	"sekolah-api/pkg/validation"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Usecase *usecase.AuthUsecase
}

func NewAuthHandler(u *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{Usecase: u}
}

// POST /login
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Validasi gagal", validation.FormatValidationError(err))
		return
	}

	res, err := h.Usecase.Login(c, req)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	response.Success(c, "Login berhasil", res)
}

// POST /refresh
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var input struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, "Validasi gagal", validation.FormatValidationError(err))
		return
	}

	newToken, err := h.Usecase.RefreshToken(input.RefreshToken)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	response.Success(c, "Token diperbarui", gin.H{
		"access_token": newToken,
	})
}

// POST /logout
func (h *AuthHandler) Logout(c *gin.Context) {
	var input struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, "Validasi gagal", validation.FormatValidationError(err))
		return
	}

	if err := h.Usecase.Logout(input.RefreshToken); err != nil {
		response.Error(c, http.StatusInternalServerError, "Gagal logout", nil)
		return
	}

	response.Success(c, "Logout berhasil", nil)
}

// POST /change-password
func (h *AuthHandler) ChangePassword(c *gin.Context) {
	var input struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, "Validasi gagal", validation.FormatValidationError(err))
		return
	}

	penggunaID, ok := c.Get("pengguna_id")
	if !ok {
		response.Error(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	err := h.Usecase.ChangePassword(c, penggunaID.(string), input.OldPassword, input.NewPassword)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response.Success(c, "Password berhasil diubah", nil)
}
