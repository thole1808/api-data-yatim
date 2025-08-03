package controllers

import (
	"api-data-yatim/config"
	"api-data-yatim/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllRT godoc
// @Summary Get semua RT
// @Description Ambil semua data RT tanpa pagination.
// @Tags RT
// @Security BearerAuth
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} models.GenericResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/rt/all [get]
func GetAllRT(c *gin.Context) {
	var rt []models.RT

	if err := config.DB.Find(&rt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to fetch RT data",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.GenericResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    rt,
	})
}

// GetRTPerPage godoc
// @Summary Get RT dengan pagination
// @Description Ambil daftar RT per halaman.
// @Tags RT
// @Security BearerAuth
// @Security ApiKeyAuth
// @Produce json
// @Param page query int false "Halaman (default 1)"
// @Param limit query int false "Jumlah data per halaman (default 10)"
// @Success 200 {object} models.GenericResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/rt [get]
func GetRTPerPage(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit

	var total int64
	var rt []models.RT

	// Hitung total data
	config.DB.Model(&models.RT{}).Count(&total)

	// Ambil data sesuai pagination
	if err := config.DB.Offset(offset).Limit(limit).Find(&rt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to fetch RT data",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.GenericResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data: gin.H{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"totalPages": (total + int64(limit) - 1) / int64(limit),
			"data":       rt,
		},
	})
}
