package controllers

import (
	"api-data-yatim/config"
	"api-data-yatim/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// // GetRT godoc
// // @Summary Get list of RT
// // @Description Ambil daftar RT, membutuhkan Bearer Token dan API_KEY
// // @Tags RT
// // @Security BearerAuth
// // @Security ApiKeyAuth
// // @Produce json
// // @Success 200 {array} models.RT
// // @Router /api/rt [get]
// func GetRT(c *gin.Context) {
// 	var rt []models.RT
// 	result := config.DB.Find(&rt)
// 	fmt.Println("DEBUG Rows Found:", result.RowsAffected)
// 	fmt.Println("DEBUG Error:", result.Error)

// 	if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, rt)
// }

// // GetAllRT godoc
// // @Summary Get RT dengan pagination
// // @Description Ambil daftar RT per halaman.
// // @Tags RT
// // @Security BearerAuth
// // @Security ApiKeyAuth
// // @Produce json
// // @Param page query int false "Halaman (default 1)"
// // @Param limit query int false "Jumlah data per halaman (default 10)"
// // @Success 200 {object} map[string]interface{}
// // @Router /api/rt [get]
// func GetAllRT(c *gin.Context) {
// 	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
// 	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
// 	if page < 1 {
// 		page = 1
// 	}
// 	offset := (page - 1) * limit

// 	var total int64
// 	var rt []models.RT

// 	// Hitung total data
// 	config.DB.Model(&models.RT{}).Count(&total)

// 	// Ambil data sesuai pagination
// 	if err := config.DB.Offset(offset).Limit(limit).Find(&rt).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Response JSON
// 	c.JSON(http.StatusOK, gin.H{
// 		"page":       page,
// 		"limit":      limit,
// 		"total":      total,
// 		"totalPages": (total + int64(limit) - 1) / int64(limit), // pembulatan ke atas
// 		"data":       rt,
// 	})
// }

// GetAllRT godoc
// @Summary Get semua RT
// @Description Ambil semua data RT tanpa pagination.
// @Tags RT
// @Security BearerAuth
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} models.RT
// @Router /api/rt/all [get]
func GetAllRT(c *gin.Context) {
	var rt []models.RT
	if err := config.DB.Find(&rt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rt)
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
// @Success 200 {object} map[string]interface{}
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": (total + int64(limit) - 1) / int64(limit),
		"data":       rt,
	})
}
