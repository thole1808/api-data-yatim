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

// GetAllStrukturOrganisasi godoc
// @Summary Ambil semua anggota Struktur Organisasi
// @Description Ambil semua anggota Struktur Organisasi beserta jabatan dan URL foto (proxy)
// @Tags StrukturOrganisasi
// @Security BearerAuth
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} models.GenericResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/struktur-organisasi/all [get]
func GetAllStrukturOrganisasi(c *gin.Context) {
	var struktur []models.StrukturOrganisasi

	if err := config.DB.
		Preload("Jabatan"). // pastikan relasi di model
		Order("urutan asc").
		Find(&struktur).Error; err != nil {

		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to fetch Struktur Organisasi data",
			Error:   err.Error(),
		})
		return
	}

	// Ambil base URL API Golang untuk proxy foto
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	// Update path foto menjadi URL proxy
	for i := range struktur {
		if struktur[i].Foto != nil && *struktur[i].Foto != "" {
			fileParts := strings.Split(*struktur[i].Foto, "/")
			filename := fileParts[len(fileParts)-1]
			struktur[i].Foto = ptr(baseURL + "/api/struktur-organisasi/foto/" + filename)
		}
	}

	c.JSON(http.StatusOK, models.GenericResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    struktur,
	})
}

// ProxyStrukturOrganisasiFoto godoc
// @Summary Proxy foto anggota Struktur Organisasi
// @Description Menampilkan foto anggota melalui Golang (proxy Laravel atau storage)
// @Tags StrukturOrganisasi
// @Produce image/png
// @Param filename path string true "Nama file foto"
// @Success 200 {file} file
// @Failure 404 {object} gin.H
// @Router /api/struktur-organisasi/foto/{filename} [get]
func ProxyStrukturOrganisasiFoto(c *gin.Context) {
	filename := c.Param("filename")

	// Path penyimpanan foto di Laravel atau server storage
	laravelBase := os.Getenv("LARAVEL_BASE_URL")
	if laravelBase == "" {
		laravelBase = "http://localhost:8000"
	}

	imageURL := laravelBase + "/storage/uploads/struktur_organisasi/" + filename

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

	c.Header("Content-Type", resp.Header.Get("Content-Type"))
	_, _ = io.Copy(c.Writer, resp.Body)
}

// Helper pointer string
func ptr(s string) *string {
	return &s
}
