// @title API Data Yatim
// @version 1.0
// @description API untuk mengelola data RT, RW, dan Pendidikan menggunakan API_KEY & Bearer Token.
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name API_KEY

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"api-data-yatim/config"
	"api-data-yatim/models"
	"api-data-yatim/routes"
	"log"

	"github.com/joho/godotenv"

	_ "api-data-yatim/docs"
)

func main() {
	_ = godotenv.Load()
	config.InitDB()

	// Auto migrate
	config.DB.AutoMigrate(&models.RT{}, &models.RW{}, &models.Pendidikan{})

	r := routes.SetupRouter()
	log.Fatal(r.Run(":8080"))
}
