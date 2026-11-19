package main

import (
	"diet-tracker-backend/config"
	"diet-tracker-backend/handlers"
	"diet-tracker-backend/middleware"
	"diet-tracker-backend/repositories"
	"diet-tracker-backend/services"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Muat file .env & Inisialisasi Database PostgreSQL (+ GORM AutoMigrate)
	db := config.ConnectDatabase()

	// 2. Dependency Injection (Merangkai Repository -> Service -> Handler)
	dietRepo := repositories.NewDietRepository(db)
	dietService := services.NewDietService(dietRepo)
	dietHandler := handlers.NewDietHandler(dietService)

	// 3. Inisialisasi Router menggunakan Framework Gin
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// 4. Grouping Endpoint REST API (/api/...)
	api := r.Group("/api")
	{
		// Endpoint Rekap & Pengaturan Target
		api.GET("/rekap", dietHandler.GetRekap)
		api.PUT("/target", dietHandler.UpdateTarget)

		// Endpoint Kamus Makanan
		api.GET("/kamus-makanan", dietHandler.GetKamusMakanan)
		api.POST("/kamus-makanan", dietHandler.TambahKamusMakanan)

		// Endpoint Catatan Makanan (CRUD)
		api.GET("/makanan", dietHandler.GetMakanan)
		api.POST("/makanan", dietHandler.CreateMakanan)
		api.DELETE("/makanan/:id", dietHandler.DeleteMakanan)

		// Endpoint Catatan Olahraga (CRUD)
		api.GET("/olahraga", dietHandler.GetOlahraga)
		api.POST("/olahraga", dietHandler.CreateOlahraga)
		api.DELETE("/olahraga/:id", dietHandler.DeleteOlahraga)
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("===============================================================")
	fmt.Println("🚀 Server Backend Go (Gin Framework + GORM PostgreSQL + .env)")
	fmt.Printf("🌐 Base Endpoint : http://localhost:%s/api\n", port)
	fmt.Println("===============================================================")

	// 5. Jalankan server sesuai APP_PORT di .env
	r.Run(":" + port)
}
