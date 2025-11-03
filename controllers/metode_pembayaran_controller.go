package controllers

import (
	"api-data-yatim/config"
	"api-data-yatim/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllMetodePembayaran godoc
// @Summary Get semua Metode Pembayaran
// @Description Ambil semua data metode pembayaran beserta jenis pembayarannya tanpa pagination.
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

	c.JSON(http.StatusOK, models.GenericResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    metode,
	})
}
