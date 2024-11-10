package repositories

import "latian-clean-architecture/models"

// BukuService adalah interface yang mendefinisikan operasi CRUD untuk buku
type BukuRepository interface {
	GetAllBuku() ([]models.Buku, error)       // Mengambil semua buku
	AddBuku(buku models.Buku) error           // Menambahkan buku baru
	GetBukuById(id int) (*models.Buku, error) // Mengambil buku berdasarkan ID
	UpdateBuku(buku models.Buku) error        // Memperbarui data buku
	DeleteBuku(id int) error                  // Menghapus buku berdasarkan ID
}
