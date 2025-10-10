package controllers

import (
	"api-data-yatim/config"
	"api-data-yatim/models"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// GetAllActivity godoc
// @Summary Get semua data aktivitas
// @Description Ambil semua data aktivitas beserta gambar-gambarnya
// @Tags Aktivitas
// @Security BearerAuth
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} models.GenericResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/aktivitas/all [get]
func GetAllActivity(c *gin.Context) {
	var activities []models.Activity

	if err := config.DB.Preload("Images").Find(&activities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to fetch aktivitas data",
			Error:   err.Error(),
		})
		return
	}

	// Ambil BASE_URL dari .env atau default ke localhost:8080
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	// Lengkapi path gambar jadi URL penuh
	for i := range activities {
		for j := range activities[i].Images {
			img := &activities[i].Images[j]
			if img.Path != "" && img.Path[0] != 'h' { // hindari path yang sudah absolute (http)
				img.Path = fmt.Sprintf("%s/%s", baseURL, img.Path)
			}
		}
	}

	c.JSON(http.StatusOK, models.GenericResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    activities,
	})
}
