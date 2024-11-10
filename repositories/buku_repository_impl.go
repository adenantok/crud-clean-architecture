package repositories

import (
	"errors"
	"latian-clean-architecture/models"

	"gorm.io/gorm"
)

// bukuService adalah struct yang mengimplementasikan BukuService
type BukuRepositoryImpl struct {
	DB *gorm.DB
}

func NewBukuRepository(db *gorm.DB) BukuRepository {
	return &BukuRepositoryImpl{DB: db}
}

func (r *BukuRepositoryImpl) GetBukuById(id int) (*models.Buku, error) {
	var buku models.Buku
	if err := r.DB.First(&buku, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("buku tidak ditemukan")
		}
		return nil, err
	}
	return &buku, nil
}

func (r *BukuRepositoryImpl) GetAllBuku() ([]models.Buku, error) {
	var bukuList []models.Buku
	// Menggunakan GORM untuk mengambil semua data buku
	if err := r.DB.Find(&bukuList).Error; err != nil {
		return nil, err
	}
	return bukuList, nil
}

func (r *BukuRepositoryImpl) AddBuku(buku models.Buku) error {
	if err := r.DB.Create(&buku).Error; err != nil {
		return err
	}
	return nil
}

func (r *BukuRepositoryImpl) UpdateBuku(buku models.Buku) error {
	if err := r.DB.Save(&buku).Error; err != nil {
		return err
	}
	return nil
}

func (r *BukuRepositoryImpl) DeleteBuku(id int) error {
	var buku models.Buku
	if err := r.DB.First(&buku, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("buku tidak ditemukan")
		}
	}
	if err := r.DB.Delete(&buku, id).Error; err != nil {
		return err
	}
	return nil
}
