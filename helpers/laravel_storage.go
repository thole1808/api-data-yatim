package helpers

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// SaveToLaravelStorage menyimpan file langsung ke Laravel storage/app/public
// Mengembalikan relative path yang disimpan di database (misal: uploads/mitra/logo/abc.jpg)
func SaveToLaravelStorage(file *multipart.FileHeader, subPath string) (string, error) {
	// Ambil Laravel project path dari environment
	laravelPath := os.Getenv("LARAVEL_STORAGE_PATH")
	if laravelPath == "" {
		// Default path jika tidak ada di env
		laravelPath = "../yatim-app/storage/app/public"
	}

	// Buat direktori lengkap
	fullDir := filepath.Join(laravelPath, subPath)
	if err := os.MkdirAll(fullDir, 0755); err != nil {
		return "", fmt.Errorf("gagal membuat direktori: %w", err)
	}

	// Generate nama file unik
	ext := filepath.Ext(file.Filename)
	randomName := RandomString(40) + ext
	fullPath := filepath.Join(fullDir, randomName)

	// Buka file upload
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("gagal membuka file upload: %w", err)
	}
	defer src.Close()

	// Buat file baru di Laravel storage
	dst, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("gagal membuat file di Laravel storage: %w", err)
	}
	defer dst.Close()

	// Copy isi file
	if _, err = io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("gagal menyalin file: %w", err)
	}

	// Return relative path untuk disimpan di database
	relativePath := filepath.Join(subPath, randomName)
	return relativePath, nil
}
