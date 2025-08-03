package controllers

import (
	"api-data-yatim/config"
	"api-data-yatim/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// LoginRequest untuk menerima input login
type LoginRequest struct {
	Username string `json:"username"` // bisa username atau email
	Password string `json:"password"`
}

// Login godoc
// @Summary Login untuk mendapatkan JWT Token
// @Description Login menggunakan username/email untuk mendapatkan token JWT.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login Request"
// @Success 200 {object} map[string]string
// @Router /login [post]
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	// Cari user berdasarkan username atau email (termasuk admin-api-saja yang baru disisipkan)
	if err := config.DB.Where("username = ?", req.Username).
		Or("email = ?", req.Username).
		First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username/email or password"})
		return
	}

	// Cek password bcrypt (Laravel & seedAdminAPI keduanya pakai bcrypt)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username/email or password"})
		return
	}

	// Buat JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	c.JSON(http.StatusOK, gin.H{
		"token":    tokenString,
		"username": user.Username,
		"email":    user.Email,
		"role":     user.Role,
	})
}
