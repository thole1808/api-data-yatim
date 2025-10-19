package routes

import (
	"api-data-yatim/controllers"
	"api-data-yatim/middlewares"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Ambil allowed origins dari env
	// originsEnv := os.Getenv("URL_API_PROXY")
	// // misal: "http://localhost:5173,http://localhost:5174,http://localhost:5222"
	// allowOrigins := strings.Split(originsEnv, ",")

	// ⚡ CORS Middleware
	r.Use(cors.New(cors.Config{

		AllowOrigins:     []string{os.Getenv("FRONTEND_URL")},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "API_KEY"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Endpoint login (tanpa API_KEY dan Bearer)
	r.POST("/login", controllers.Login)

	// ✅ Route publik untuk gambar — tidak butuh JWT/API key
	r.GET("/api/galeri/:filename", controllers.ProxyLaravelImage)

	r.GET("/api/struktur-organisasi/foto/:filename", controllers.ProxyStrukturOrganisasiFoto)

	r.GET("/api/mitra/logo/:filename", controllers.ProxyMitraLogo)

	api := r.Group("/api", middlewares.APIKeyMiddleware(), middlewares.AuthMiddleware())
	{
		api.GET("/rt", controllers.GetRTPerPage)
		api.GET("/rt/all", controllers.GetAllRT)
		// api.GET("/rw", controllers.GetRW)
		// api.GET("/pendidikan", controllers.GetPendidikan)

		// ===== Profil Yayasan =====
		api.GET("/profil-yayasan/all", controllers.GetAllYayasan)

		// ===== Aktivitas =====
		api.GET("/aktivitas/all", controllers.GetAllActivity)

		// ===== Struktur Organisasi =====
		api.GET("/struktur-organisasi/all", controllers.GetAllStrukturOrganisasi)
		// api.GET("/struktur-organisasi/foto/:filename", controllers.ProxyStrukturOrganisasiFoto)

		// ===== Mita =====
		api.GET("/mitra/all", controllers.GetAllMitra)
		// api.GET("/struktur-organisasi/foto/:filename", controllers.ProxyStrukturOrganisasiFoto)

	}

	return r
}
