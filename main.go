package main

import (
	"latian-clean-architecture/config"
	"latian-clean-architecture/controllers"
	"latian-clean-architecture/repositories"
	"latian-clean-architecture/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Membuat koneksi database
	db := config.ConnectDB()

	// 2. Membuat instance repository
	repo := repositories.NewBukuRepository(db)

	// 3. Membuat instance service dengan meng-inject 'repo'
	service := services.NewBukuService(repo)

	// 4. Membuat instance controller dengan meng-inject 'service'
	controller := controllers.NewBukuController(service)

	// 5. Inisialisasi router
	router := gin.Default()
	router.GET("/buku/:id", controller.GetBukuByIdHandler)
	router.GET("/buku", controller.GetAllBukuHandler)
	router.POST("/buku", controller.AddBukuHandler)
	router.PUT("/update", controller.UpdateBuku)
	router.DELETE("/hapus/:id", controller.DeleteBukuHandler)
	router.Run(":3000")
}
