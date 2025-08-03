package routes

import (
	"api-data-yatim/controllers"
	"api-data-yatim/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Endpoint login (tanpa API_KEY dan Bearer)
	r.POST("/login", controllers.Login)

	api := r.Group("/api", middlewares.APIKeyMiddleware(), middlewares.AuthMiddleware())
	{
		api.GET("/rt", controllers.GetRTPerPage) // ✅ RT per halaman
		api.GET("/rt/all", controllers.GetAllRT) // ✅ RT semua data
		// api.GET("/rw", controllers.GetRW)
		// api.GET("/pendidikan", controllers.GetPendidikan)
	}

	return r
}
