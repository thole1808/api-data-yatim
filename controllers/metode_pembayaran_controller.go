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

// ==============================
// 📘 GET: Semua Metode Pembayaran
// ==============================

// GetAllMetodePembayaran godoc
// @Summary Ambil semua Metode Pembayaran
// @Description Ambil semua data metode pembayaran beserta jenis pembayaran dan URL QR (proxy)
// @Tags Metode Pembayaran
// @Security BearerAuth
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} models.GenericResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/metode-pembayaran/all [get]
func GetAllMetodePembayaran(c *gin.Context) {
	var metode []models.MetodePembayaran

	if err := config.DB.Preload("JenisPembayaran").Find(&metode).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to fetch metode pembayaran data",
			Error:   err.Error(),
		})
		return
	}

	// 🔹 Base URL proxy untuk QR
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	// 🔹 Update path QR menjadi URL proxy
	for i := range metode {
		if metode[i].QRImage != "" {
			fileParts := strings.Split(metode[i].QRImage, "/")
			filename := fileParts[len(fileParts)-1]
			metode[i].QRImage = baseURL + "/api/metode-pembayaran/qr/" + filename
		}
	}

	c.JSON(http.StatusOK, models.GenericResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    metode,
	})
}

// ==============================
// 🖼️ GET: Proxy QR Metode Pembayaran
// ==============================

// ProxyMetodePembayaranQR godoc
// @Summary Proxy QR image metode pembayaran
// @Description Menampilkan QR image metode pembayaran melalui Golang (proxy dari Laravel storage)
// @Tags Metode Pembayaran
// @Produce image/png
// @Param filename path string true "Nama file QR"
// @Success 200 {file} file
// @Failure 404 {object} models.ErrorResponse
// @Router /api/metode-pembayaran/qr/{filename} [get]
func ProxyMetodePembayaranQR(c *gin.Context) {
	filename := c.Param("filename")

	// 🔹 Base URL Laravel
	laravelBase := os.Getenv("LARAVEL_BASE_URL")
	if laravelBase == "" {
		laravelBase = "http://localhost:8000"
	}

	imageURL := laravelBase + "/storage/uploads/metode_pembayaran/" + filename

	// 🔹 Request gambar dari Laravel
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

	// 🔹 Return langsung stream gambar
	c.Header("Content-Type", resp.Header.Get("Content-Type"))
	_, _ = io.Copy(c.Writer, resp.Body)
}
