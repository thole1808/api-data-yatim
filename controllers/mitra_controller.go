package controllers

import (
	"api-data-yatim/config"
	"api-data-yatim/helpers" // 🟢 tambahkan helper global untuk pointer
	"api-data-yatim/models"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// ==============================
// 📘 GET: Semua Mitra
// ==============================

// GetAllMitra godoc
// @Summary Ambil semua Mitra
// @Description Mengambil semua data Mitra beserta kategori_mitra dan URL logo (proxy dari Laravel storage)
// @Tags Mitra
// @Security BearerAuth
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} models.GenericResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/mitra/all [get]
func GetAllMitra(c *gin.Context) {
	var mitra []models.Mitra

	// 🔹 Ambil data mitra lengkap beserta relasi kategori
	if err := config.DB.
		Preload("KategoriMitra").
		Order("id ASC").
		Find(&mitra).Error; err != nil {

		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to fetch Mitra data",
			Error:   err.Error(),
		})
		return
	}

	// 🔹 Base URL proxy (untuk logo)
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	// 🔹 Ubah path logo menjadi URL proxy
	for i := range mitra {
		if mitra[i].Logo != nil && *mitra[i].Logo != "" {
			fileParts := strings.Split(*mitra[i].Logo, "/")
			filename := fileParts[len(fileParts)-1]
			mitra[i].Logo = helpers.Ptr(baseURL + "/api/mitra/logo/" + filename)
		}
	}

	c.JSON(http.StatusOK, models.GenericResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    mitra,
	})
}

// ==============================
// 🖼️ GET: Proxy Logo Mitra
// ==============================

// ProxyMitraLogo godoc
// @Summary Proxy logo mitra
// @Description Menampilkan logo mitra melalui Golang (proxy dari Laravel storage)
// @Tags Mitra
// @Produce image/png
// @Param filename path string true "Nama file logo"
// @Success 200 {file} file
// @Failure 404 {object} models.ErrorResponse
// @Router /api/mitra/logo/{filename} [get]
func ProxyMitraLogo(c *gin.Context) {
	filename := c.Param("filename")

	// 🔹 Base URL Laravel
	laravelBase := os.Getenv("LARAVEL_BASE_URL")
	if laravelBase == "" {
		laravelBase = "http://localhost:8000"
	}

	imageURL := laravelBase + "/storage/uploads/mitra/logo/" + filename

	// 🔹 Ambil gambar dari Laravel
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

	// 🔹 Return stream gambar langsung
	c.Header("Content-Type", resp.Header.Get("Content-Type"))
	_, _ = io.Copy(c.Writer, resp.Body)
}

// ==============================
// 💚 POST: Tambah Mitra Personal (Donasi Pribadi)
// ==============================

// AddMitraPersonal godoc
// @Summary Tambah Donasi Personal
// @Description Menambahkan data donasi personal (tanpa relasi langsung ke mitra lain). Data dikirim via form-data beserta bukti transfer.
// @Tags Donasi
// @Accept multipart/form-data
// @Security BearerAuth
// @Security ApiKeyAuth
// @Produce json
// @Param nama formData string true "Nama Donatur"
// @Param jumlah formData string true "Jumlah Donasi (Rp)"
// @Param pesan formData string false "Pesan Donasi"
// @Param metode formData string false "Metode Pembayaran (qris/bank)"
// @Param bukti formData file false "Upload Bukti Transfer"
// @Success 200 {object} models.Mitra
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/mitra/personal [post]
func AddMitraPersonal(c *gin.Context) {
	var input struct {
		Nama   string `form:"nama" binding:"required"`
		Jumlah string `form:"jumlah" binding:"required"`
		Pesan  string `form:"pesan"`
		Metode string `form:"metode"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data form tidak lengkap"})
		return
	}

	// 🔹 Ambil file bukti upload (opsional)
	file, err := c.FormFile("bukti")
	var filePath *string
	if err == nil {
		uploadDir := "storage/uploads/mitra"
		os.MkdirAll(uploadDir, os.ModePerm)

		filename := strings.ReplaceAll(input.Nama, " ", "_") + "_" + file.Filename
		fullPath := filepath.Join(uploadDir, filename)

		if err := c.SaveUploadedFile(file, fullPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan bukti transfer"})
			return
		}
		filePath = &fullPath
	}

	// 🔹 Pastikan kategori personal ada
	var kategori models.KategoriMitra
	if err := config.DB.First(&kategori, 3).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Kategori 'Personal' belum tersedia di database",
		})
		return
	}

	// 🔹 Simpan data ke tabel mitra
	mitra := models.Mitra{
		Nama:            input.Nama,
		Deskripsi:       helpers.Ptr(input.Pesan),
		Logo:            filePath,
		KategoriMitraID: &kategori.ID,
	}

	if err := config.DB.Create(&mitra).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data ke database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Donasi personal berhasil disimpan",
		"data":    mitra,
	})
}
