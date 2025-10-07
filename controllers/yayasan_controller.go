package controllers

import (
	"api-data-yatim/config"
	"api-data-yatim/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllYayasan godoc
// @Summary Get semua Profil Yayasan
// @Description Ambil semua data profil yayasan tanpa pagination.
// @Tags Profil Yayasan
// @Security BearerAuth
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} models.GenericResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/profil-yayasan/all [get]
func GetAllYayasan(c *gin.Context) {
	var yayasan []models.ProfilYayasan

	if err := config.DB.Find(&yayasan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to fetch profil yayasan data",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.GenericResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    yayasan,
	})
}
