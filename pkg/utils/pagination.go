package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPaginationParams(c *gin.Context) (int, int) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	return page, limit
}

func CalculateTotalPages(totalItems, limit int) int {
	if limit == 0 {
		return 0
	}
	pages := totalItems / limit
	if totalItems%limit != 0 {
		pages++
	}
	return pages
}
