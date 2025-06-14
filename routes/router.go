package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"sekolah-api/internal/middleware"
	userHttp "sekolah-api/internal/pengguna/delivery/http"
	userRepo "sekolah-api/internal/pengguna/infrastructure/persistence"
	userUse "sekolah-api/internal/pengguna/usecase"

	authHttp "sekolah-api/internal/auth/delivery/http"
	authRepo "sekolah-api/internal/auth/infrastructure/persistence"
	authUse "sekolah-api/internal/auth/usecase"

	siswaHttp "sekolah-api/internal/siswa/delivery/http"
	siswaRepo "sekolah-api/internal/siswa/infrastructure/persistence"
	siswaUse "sekolah-api/internal/siswa/usecase"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.Logger())

	api := r.Group("/api/v1")

	// Inisialisasi dependency
	uRepo := userRepo.NewPenggunaRepository(db)
	uUse := userUse.NewPenggunaUsecase(uRepo)
	uHandler := userHttp.NewPenggunaHandler(uUse)

	auRepo := authRepo.NewRefreshTokenRepo(db)
	auUse := authUse.NewAuthUsecase(uRepo, auRepo)
	auHandler := authHttp.NewAuthHandler(auUse)

	siRepo := siswaRepo.NewSiswaRepository(db)
	siUse := siswaUse.NewSiswaUsecase(siRepo)
	siHandler := siswaHttp.NewSiswaHandler(siUse)

	// Grup Auth
	auth := api.Group("/auth")
	{
		auth.POST("/login", auHandler.Login)
		auth.POST("/refresh", auHandler.RefreshToken)
		auth.POST("/logout", middleware.Auth(), auHandler.Logout)
		auth.POST("/change-password", middleware.Auth(), auHandler.ChangePassword)
	}

	// Grup Pengguna dengan middleware
	pengguna := api.Group("/pengguna").Use(middleware.Auth())
	{
		pengguna.GET("/", uHandler.GetAll)
		pengguna.GET("/:id", uHandler.GetByID)
		pengguna.POST("/", uHandler.Create)
		pengguna.PUT("/:id", uHandler.Update)
		pengguna.DELETE("/:id", uHandler.Delete)
	}

	siswa := api.Group("/siswa").Use(middleware.Auth())
	{
		siswa.GET("/", siHandler.GetAll)
		siswa.GET("/:id", siHandler.GetByID)
		siswa.POST("/", siHandler.Create)
		siswa.PUT("/:id", siHandler.Update)
		siswa.DELETE("/:id", siHandler.Delete)
	}

	return r
}
