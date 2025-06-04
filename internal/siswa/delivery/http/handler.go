package http

import (
	"net/http"
	"sekolah-api/internal/siswa/dto"
	"sekolah-api/internal/siswa/usecase"
	"sekolah-api/pkg/response"
	"sekolah-api/pkg/utils"
	"sekolah-api/pkg/validation"

	"github.com/gin-gonic/gin"
)

type SiswaHandler struct {
	Usecase *usecase.SiswaUsecase
}

func NewSiswaHandler(u *usecase.SiswaUsecase) *SiswaHandler {
	return &SiswaHandler{Usecase: u}
}

func (h *SiswaHandler) Create(c *gin.Context) {
	var req dto.CreateSiswaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Validasi gagal", validation.FormatValidationError(err))
		return
	}

	if err := h.Usecase.Create(c.Request.Context(), req); err != nil {
		response.Error(c, http.StatusInternalServerError, "Gagal membuat siswa", nil)
		return
	}

	response.Success(c, "Siswa berhasil dibuat", nil)
}

func (h *SiswaHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateSiswaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Validasi gagal", validation.FormatValidationError(err))
		return
	}

	if err := h.Usecase.Update(c.Request.Context(), id, req); err != nil {
		response.Error(c, http.StatusInternalServerError, "Gagal memperbarui siswa", nil)
		return
	}

	response.Success(c, "Siswa diperbarui", nil)
}

func (h *SiswaHandler) GetAll(c *gin.Context) {
	page, limit := utils.GetPaginationParams(c)

	result, total, err := h.Usecase.GetAll(c.Request.Context(), page, limit)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Gagal mengambil data", nil)
		return
	}

	pagination := response.Pagination{
		Page:       page,
		Limit:      limit,
		TotalItems: total,
		TotalPages: utils.CalculateTotalPages(total, limit),
	}

	response.Paginated(c, "Data siswa", result, pagination)
}

func (h *SiswaHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	result, err := h.Usecase.GetByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "Siswa tidak ditemukan", nil)
		return
	}

	response.Success(c, "Detail siswa", result)
}

func (h *SiswaHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Usecase.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, "Gagal menghapus siswa", nil)
		return
	}

	response.Success(c, "Siswa dihapus", nil)
}
