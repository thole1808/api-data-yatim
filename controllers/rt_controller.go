package controllers

import (
	"api-data-yatim/config"
	"api-data-yatim/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRT godoc
// @Summary Get list of RT
// @Description Ambil daftar RT, membutuhkan Bearer Token dan API_KEY
// @Tags RT
// @Security BearerAuth
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} models.RT
// @Router /api/rt [get]
func GetRT(c *gin.Context) {
	var rt []models.RT
	result := config.DB.Find(&rt)
	fmt.Println("DEBUG Rows Found:", result.RowsAffected)
	fmt.Println("DEBUG Error:", result.Error)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, rt)
}
