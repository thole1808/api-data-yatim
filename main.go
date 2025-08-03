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
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"

	_ "api-data-yatim/docs"
)

func main() {
	_ = godotenv.Load()
	config.InitDB()
	seedAdminAPI()

	// Auto migrate
	// config.DB.AutoMigrate(&models.User{}, &models.RT{}, &models.RW{}, &models.Pendidikan{})
	config.DB.AutoMigrate(&models.RT{}, &models.RW{}, &models.Pendidikan{})

	// Seeder user admin
	// hashedPass, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	// config.DB.FirstOrCreate(&models.User{}, models.User{Username: "adminapi", Password: string(hashedPass)})

	r := routes.SetupRouter()
	log.Fatal(r.Run(":8080"))
}

func seedAdminAPI() {
	var user models.User
	if err := config.DB.Where("username = ?", "admin-api-saja").First(&user).Error; err != nil {
		hashedPass, _ := bcrypt.GenerateFromPassword([]byte("adminapi123"), bcrypt.DefaultCost)
		config.DB.Exec(`
            INSERT INTO users (name, email, username, password, role, created_at, updated_at)
            VALUES (?, ?, ?, ?, ?, NOW(), NOW())
        `, "Admin API", "adminapi@example.com", "admin-api-saja", string(hashedPass), "admin-api")
		fmt.Println("✅ User admin-api-saja berhasil dibuat")
	}
}
