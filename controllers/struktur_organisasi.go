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
// @Description Ambil semua anggota Struktur Organisasi, beserta nama jabatan dari master jabatan dan URL foto
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

	// Ambil base URL API untuk foto
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	// Ubah path foto menjadi URL lengkap
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

// Helper untuk pointer string
func ptr(s string) *string {
	return &s
}

// ProxyStrukturOrganisasiFoto -> endpoint proxy foto anggota
func ProxyStrukturOrganisasiFoto(c *gin.Context) {
	filename := c.Param("filename")
	// Path penyimpanan di Laravel atau server storage
	imagePath := "storage/uploads/struktur_organisasi/" + filename

	file, err := os.Open(imagePath)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}
	defer file.Close()

	// Set header sesuai tipe file
	c.Header("Content-Type", "image/png")
	_, _ = io.Copy(c.Writer, file)
}
