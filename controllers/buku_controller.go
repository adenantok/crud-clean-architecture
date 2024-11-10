package controllers

import (
	"latian-clean-architecture/models"
	"latian-clean-architecture/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BukuController struct {
	service *services.BukuService
}

func NewBukuController(service *services.BukuService) *BukuController {
	return &BukuController{service: service}
}

// GetBukuByIdHandler menangani permintaan GET berdasarkan ID buku
func (bc *BukuController) GetBukuByIdHandler(c *gin.Context) {
	id := c.Param("id")

	// Konversi ID dari string ke int
	bukuId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	// Panggil service untuk mendapatkan buku berdasarkan ID
	buku, err := bc.service.GetBukuById(bukuId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "buku tidak ditemukan"})
		return
	}

	// Berhasil mendapatkan buku berdasarkan ID
	c.JSON(http.StatusOK, gin.H{"data": buku})
}

// GetBukuHandler menangani permintaan GET untuk mendapatkan semua data buku
func (bc *BukuController) GetAllBukuHandler(c *gin.Context) {
	buku, err := bc.service.GetAllBuku()
	if err != nil {
		// Jika terjadi kesalahan saat mengambil data, kirim respons error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data buku"})
		return
	}
	// Berhasil mendapatkan data buku
	c.JSON(http.StatusOK, gin.H{"data": buku})
}

// AddBukuHandler menangani permintaan POST untuk menambah buku
func (bc *BukuController) AddBukuHandler(c *gin.Context) {
	var buku models.Buku

	// Bind JSON dari request ke struct buku
	if err := c.ShouldBindJSON(&buku); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	// Validasi: Pastikan 'judul' tidak kosong dan 'harga' ada
	if buku.Judul == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Field 'judul' tidak boleh kosong"})
		return
	}
	if buku.Harga == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Field 'harga' tidak boleh kosong"})
		return
	}

	// Panggil service untuk menambahkan buku ke database
	if err := bc.service.AddBuku(buku); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambah data buku"})
		return
	}

	// Berhasil menambah buku
	c.JSON(http.StatusOK, gin.H{"message": "Data buku berhasil ditambahkan"})
}

// UpdateBuku menangani permintaan PUT untuk memperbarui buku
func (bc *BukuController) UpdateBuku(c *gin.Context) {
	var buku models.Buku
	if err := c.ShouldBindJSON(&buku); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	// Validasi input: cek apakah 'id', 'judul', dan 'harga' tidak kosong
	var validationErrors []string
	if buku.Id == nil {
		validationErrors = append(validationErrors, "Field 'id' tidak boleh kosong")
	}
	if buku.Judul == "" {
		validationErrors = append(validationErrors, "Field 'judul' tidak boleh kosong")
	}
	if buku.Harga == nil {
		validationErrors = append(validationErrors, "Field 'harga' tidak boleh kosong")
	}

	if len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	if err := bc.service.UpdateBuku(buku); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui data buku"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data buku berhasil diperbarui"})
}

// DeleteBukuHandler menangani permintaan DELETE berdasarkan ID buku
func (bc *BukuController) DeleteBukuHandler(c *gin.Context) {
	id := c.Param("id")
	bukuId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	err = bc.service.DeleteBuku(bukuId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Buku dengan ID " + id + " berhasil dihapus"})
}
