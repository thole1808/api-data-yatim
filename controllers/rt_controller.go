package controllers

import (
	"github.com/gin-gonic/gin"
)

// // @Summary Get all RT
// // @Tags RT
// // @Produce json
// // @Security ApiKeyAuth
// // @Security BearerAuth
// // @Success 200 {array} models.RT
// // @Router /api/rt [get]
// func GetRT(c *gin.Context) {
// 	var rt []models.RT
// 	config.DB.Find(&rt)
// 	c.JSON(http.StatusOK, rt)
// }

// GetRT godoc
// @Summary Get all RT
// @Description Ambil semua data RT
// @Tags RT
// @Produce json
// @Security ApiKeyAuth
// @Security BearerAuth
// @Success 200 {array} models.RT
// @Router /api/rt [get]
func GetRT(c *gin.Context) {
}
