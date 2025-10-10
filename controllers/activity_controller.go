// // package controllers

// // import (
// // 	"api-data-yatim/config"
// // 	"api-data-yatim/models"
// // 	"net/http"
// // 	"os"
// // 	"strings"

// // 	"github.com/gin-gonic/gin"
// // )

// // // GetAllActivity godoc
// // // @Summary Get semua data aktivitas
// // // @Description Ambil semua data aktivitas beserta gambar-gambarnya
// // // @Tags Aktivitas
// // // @Security BearerAuth
// // // @Security ApiKeyAuth
// // // @Produce json
// // // @Success 200 {object} models.GenericResponse
// // // @Failure 500 {object} models.ErrorResponse
// // // @Router /api/aktivitas/all [get]
// // func GetAllActivity(c *gin.Context) {
// // 	var activities []models.Activity

// // 	if err := config.DB.Preload("Images").Find(&activities).Error; err != nil {
// // 		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
// // 			Code:    http.StatusInternalServerError,
// // 			Message: "Failed to fetch aktivitas data",
// // 			Error:   err.Error(),
// // 		})
// // 		return
// // 	}

// // 	// Ambil base URL Laravel dari .env
// // 	laravelBase := os.Getenv("LARAVEL_BASE_URL")
// // 	if laravelBase == "" {
// // 		laravelBase = "http://localhost:8000/storage" // default jika tidak ada di .env
// // 	}

// // 	// Ubah path gambar agar jadi URL lengkap
// // 	for i := range activities {
// // 		for j := range activities[i].Images {
// // 			path := activities[i].Images[j].Path
// // 			if !strings.HasPrefix(path, "http") {
// // 				activities[i].Images[j].Path = laravelBase + "/" + path
// // 			}
// // 		}
// // 	}

// // 	c.JSON(http.StatusOK, models.GenericResponse{
// // 		Code:    http.StatusOK,
// // 		Message: "Success",
// // 		Data:    activities,
// // 	})
// // }

// package controllers

// import (
// 	"api-data-yatim/config"
// 	"api-data-yatim/models"
// 	"io"
// 	"net/http"
// 	"os"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// )

// // GetAllActivity godoc
// // @Summary Get semua data aktivitas
// // @Description Ambil semua data aktivitas beserta gambar-gambarnya
// // @Tags Aktivitas
// // @Security BearerAuth
// // @Security ApiKeyAuth
// // @Produce json
// // @Success 200 {object} models.GenericResponse
// // @Failure 500 {object} models.ErrorResponse
// // @Router /api/aktivitas/all [get]
// func GetAllActivity(c *gin.Context) {
// 	var activities []models.Activity

// 	if err := config.DB.Preload("Images").Find(&activities).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
// 			Code:    http.StatusInternalServerError,
// 			Message: "Failed to fetch aktivitas data",
// 			Error:   err.Error(),
// 		})
// 		return
// 	}

// 	// Ambil base URL Golang (bukan Laravel)
// 	baseURL := os.Getenv("BASE_URL")
// 	if baseURL == "" {
// 		baseURL = "http://localhost:8080"
// 	}

// 	// Ubah setiap path gambar agar diarahkan ke proxy Golang
// 	for i := range activities {
// 		for j := range activities[i].Images {
// 			path := activities[i].Images[j].Path

// 			// Ambil nama file terakhir (contoh: uploads/activities/foto1.jpg -> foto1.jpg)
// 			fileParts := strings.Split(path, "/")
// 			filename := fileParts[len(fileParts)-1]

// 			// Path aman via proxy Golang
// 			activities[i].Images[j].Path = baseURL + "/api/galeri/" + filename
// 		}
// 	}

// 	c.JSON(http.StatusOK, models.GenericResponse{
// 		Code:    http.StatusOK,
// 		Message: "Success",
// 		Data:    activities,
// 	})
// }

// // ProxyLaravelImage -> proxy endpoint untuk menampilkan gambar Laravel lewat Golang
// func ProxyLaravelImage(c *gin.Context) {
// 	filename := c.Param("filename")
// 	laravelBase := os.Getenv("LARAVEL_BASE_URL")
// 	if laravelBase == "" {
// 		laravelBase = "http://localhost:8000"
// 	}

// 	// Path asli di Laravel
// 	imageURL := laravelBase + "/storage/uploads/activities/" + filename

// 	resp, err := http.Get(imageURL)
// 	if err != nil {
// 		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to fetch image from Laravel"})
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		c.JSON(resp.StatusCode, gin.H{"error": "Image not found"})
// 		return
// 	}

// 	// Set header dan stream gambar
// 	c.Header("Content-Type", resp.Header.Get("Content-Type"))
// 	// Stream image data to response writer
// 	_, _ = io.Copy(c.Writer, resp.Body)
// }

package controllers

import (
	"api-data-yatim/config"
	"api-data-yatim/models"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetAllActivity godoc
// @Summary Get semua data aktivitas
// @Description Ambil semua data aktivitas beserta gambar dan nama kategori
// @Tags Aktivitas
// @Security BearerAuth
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} models.GenericResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/aktivitas/all [get]
func GetAllActivity(c *gin.Context) {
	var activities []models.Activity

	// Preload relasi Kategori dan Images
	if err := config.DB.
		Preload("Images").
		Preload("Kategori").
		Find(&activities).Error; err != nil {

		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to fetch aktivitas data",
			Error:   err.Error(),
		})
		return
	}

	// Ambil base URL API Golang
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	// Ubah setiap path gambar agar diarahkan ke proxy Golang
	for i := range activities {
		for j := range activities[i].Images {
			path := activities[i].Images[j].Path

			// Ambil nama file terakhir dari path
			fileParts := strings.Split(path, "/")
			filename := fileParts[len(fileParts)-1]

			// Buat path baru via endpoint proxy
			activities[i].Images[j].Path = baseURL + "/api/galeri/" + filename
		}
	}

	c.JSON(http.StatusOK, models.GenericResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    activities,
	})
}

// ProxyLaravelImage -> proxy endpoint untuk menampilkan gambar Laravel lewat Golang
func ProxyLaravelImage(c *gin.Context) {
	filename := c.Param("filename")
	laravelBase := os.Getenv("LARAVEL_BASE_URL")
	if laravelBase == "" {
		laravelBase = "http://localhost:8000"
	}

	// Path asli di Laravel
	imageURL := laravelBase + "/storage/uploads/activities/" + filename

	resp, err := http.Get(imageURL)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to fetch image from Laravel"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "Image not found"})
		return
	}

	// Set header konten
	c.Header("Content-Type", resp.Header.Get("Content-Type"))
	_, _ = io.Copy(c.Writer, resp.Body)
}
