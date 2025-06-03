package http

import (
	"net/http"
	"sekolah-api/internal/pengguna/dto"
	"sekolah-api/internal/pengguna/usecase"
	"sekolah-api/pkg/response"
	"sekolah-api/pkg/utils"
	"sekolah-api/pkg/validation"

	"github.com/gin-gonic/gin"
)

type PenggunaHandler struct {
	Usecase *usecase.PenggunaUsecase
}

func NewPenggunaHandler(u *usecase.PenggunaUsecase) *PenggunaHandler {
	return &PenggunaHandler{Usecase: u}
}

func (h *PenggunaHandler) Create(c *gin.Context) {
	var req dto.CreatePenggunaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Validasi gagal", validation.FormatValidationError(err))
		return
	}

	if err := h.Usecase.Create(c.Request.Context(), req); err != nil {
		response.Error(c, http.StatusInternalServerError, "Gagal membuat pengguna", nil)
		return
	}

	response.Success(c, "Pengguna berhasil dibuat", nil)
}

func (h *PenggunaHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdatePenggunaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Validasi gagal", validation.FormatValidationError(err))
		return
	}

	if err := h.Usecase.Update(c.Request.Context(), id, req); err != nil {
		response.Error(c, http.StatusInternalServerError, "Gagal memperbarui pengguna", nil)
		return
	}

	response.Success(c, "Pengguna diperbarui", nil)
}

func (h *PenggunaHandler) GetAll(c *gin.Context) {
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

	response.Paginated(c, "Data pengguna", result, pagination)
}

func (h *PenggunaHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	result, err := h.Usecase.GetByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "Pengguna tidak ditemukan", nil)
		return
	}

	response.Success(c, "Detail pengguna", result)
}

func (h *PenggunaHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Usecase.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, "Gagal menghapus pengguna", nil)
		return
	}

	response.Success(c, "Pengguna dihapus", nil)
}
