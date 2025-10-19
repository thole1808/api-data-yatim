package controllers

import (
	"api-data-yatim/config"
	"api-data-yatim/helpers" // 🟢 tambahkan helper global untuk pointer
	"api-data-yatim/models"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// ==============================
// 📘 GET: Semua Mitra
// ==============================

// GetAllMitra godoc
// @Summary Ambil semua Mitra
// @Description Mengambil semua data Mitra beserta kategori_mitra dan URL logo (proxy dari Laravel storage)
// @Tags Mitra
// @Security BearerAuth
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} models.GenericResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/mitra/all [get]
func GetAllMitra(c *gin.Context) {
	var mitra []models.Mitra

	// 🔹 Ambil data mitra lengkap beserta relasi kategori
	if err := config.DB.
		Preload("KategoriMitra").
		Order("id ASC").
		Find(&mitra).Error; err != nil {

		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to fetch Mitra data",
			Error:   err.Error(),
		})
		return
	}

	// 🔹 Base URL proxy (untuk logo)
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	// 🔹 Ubah path logo menjadi URL proxy
	for i := range mitra {
		if mitra[i].Logo != nil && *mitra[i].Logo != "" {
			fileParts := strings.Split(*mitra[i].Logo, "/")
			filename := fileParts[len(fileParts)-1]
			mitra[i].Logo = helpers.Ptr(baseURL + "/api/mitra/logo/" + filename)
		}
	}

	c.JSON(http.StatusOK, models.GenericResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    mitra,
	})
}

// ==============================
// 🖼️ GET: Proxy Logo Mitra
// ==============================

// ProxyMitraLogo godoc
// @Summary Proxy logo mitra
// @Description Menampilkan logo mitra melalui Golang (proxy dari Laravel storage)
// @Tags Mitra
// @Produce image/png
// @Param filename path string true "Nama file logo"
// @Success 200 {file} file
// @Failure 404 {object} gin.H
// @Router /api/mitra/logo/{filename} [get]
func ProxyMitraLogo(c *gin.Context) {
	filename := c.Param("filename")

	// 🔹 Base URL Laravel
	laravelBase := os.Getenv("LARAVEL_BASE_URL")
	if laravelBase == "" {
		laravelBase = "http://localhost:8000"
	}

	imageURL := laravelBase + "/storage/uploads/mitra/logo/" + filename

	// 🔹 Ambil gambar dari Laravel
	resp, err := http.Get(imageURL)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to fetch image from storage"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "Image not found"})
		return
	}

	// 🔹 Return stream gambar langsung
	c.Header("Content-Type", resp.Header.Get("Content-Type"))
	_, _ = io.Copy(c.Writer, resp.Body)
}
