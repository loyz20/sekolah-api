package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Status  string      `json:"status"`            // "success" atau "error"
	Message string      `json:"message,omitempty"` // Keterangan singkat
	Data    interface{} `json:"data,omitempty"`    // Payload jika sukses
	Errors  interface{} `json:"errors,omitempty"`  // Detail error jika gagal
}

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

type PaginatedData struct {
	Items      interface{} `json:"items"`
	Pagination Pagination  `json:"pagination"`
}

// Success response dengan data
func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// Created 201
func Created(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// Error response biasa
func Error(c *gin.Context, statusCode int, message string, err interface{}) {
	c.JSON(statusCode, APIResponse{
		Status:  "error",
		Message: message,
		Errors:  err,
	})
}

func Paginated(c *gin.Context, message string, items interface{}, pagination Pagination) {
	c.JSON(http.StatusOK, APIResponse{
		Status:  "success",
		Message: message,
		Data: PaginatedData{
			Items:      items,
			Pagination: pagination,
		},
	})
}
