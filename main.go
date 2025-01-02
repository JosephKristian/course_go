package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gitlab.com/ipaymupreviews/golang-gin-poc/config"
	"gitlab.com/ipaymupreviews/golang-gin-poc/routes"
)

func main() {
	// Menggunakan gin.Default() yang sudah menyertakan middleware default
	router := gin.Default()

	// Koneksi ke database
	config.Connect()

	// Menyusun routing
	routes.UserRoute(router)
	routes.AuthRoutes(router)

	// Menjalankan server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
