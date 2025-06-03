package main

import (
	"sekolah-api/config"
	"sekolah-api/pkg/seeder"
	"sekolah-api/routes"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	db := config.InitDB()
	seeder.SeedPengguna(db)
	r := routes.SetupRouter(db)

	r.Run(":8080")
}
