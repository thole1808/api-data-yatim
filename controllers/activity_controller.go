package controllers

import (
	"api-data-yatim/config"
	"api-data-yatim/models"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// GetAllActivity godoc
// @Summary Get semua data aktivitas
// @Description Ambil semua data aktivitas yang berstatus published dan sedang aktif (tanggal mulai <= sekarang <= tanggal selesai), beserta gambar & kategori
// @Tags Aktivitas
// @Security BearerAuth
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} models.GenericResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/aktivitas/all [get]
func GetAllActivity(c *gin.Context) {
	var activities []models.Activity

	now := time.Now() // waktu saat ini

	// Filter hanya yang status "published" dan tanggal aktif
	if err := config.DB.
		Preload("Images").
		Preload("Kategori").
		Where("status = ?", "published").
		Where("tanggal_mulai <= ?", now).
		Where("tanggal_selesai >= ?", now).
		Find(&activities).Error; err != nil {

		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to fetch aktivitas data",
			Error:   err.Error(),
		})
		return
	}

	// Ambil base URL API Golang (untuk path gambar)
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	// Ubah path gambar menjadi URL dari proxy
	for i := range activities {
		for j := range activities[i].Images {
			path := activities[i].Images[j].Path
			fileParts := strings.Split(path, "/")
			filename := fileParts[len(fileParts)-1]
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

	c.Header("Content-Type", resp.Header.Get("Content-Type"))
	_, _ = io.Copy(c.Writer, resp.Body)
}
