package middleware

import (
	"net/http"
	"sekolah-api/pkg/response"
	"sekolah-api/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			response.Error(c, http.StatusUnauthorized, "Token tidak ditemukan", nil)
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.ParseWithClaims(tokenStr, &utils.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return utils.GetSecret(), nil
		})

		if err != nil || !token.Valid {
			response.Error(c, http.StatusUnauthorized, "Token tidak valid", nil)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*utils.JWTClaims)
		if !ok {
			response.Error(c, http.StatusUnauthorized, "Token tidak valid", nil)
			c.Abort()
			return
		}

		c.Set("pengguna_id", claims.PenggunaID)
		c.Set("sekolah_id", claims.SekolahID)
		c.Set("peran_id_str", claims.PeranIDStr)

		c.Next()
	}
}
