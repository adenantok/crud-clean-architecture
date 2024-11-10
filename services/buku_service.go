package services

import (
	"latian-clean-architecture/models"
	"latian-clean-architecture/repositories"
)

type BukuService struct {
	repo repositories.BukuRepository
}

func NewBukuService(repo repositories.BukuRepository) *BukuService {
	return &BukuService{repo: repo}
}

func (bs *BukuService) GetBukuById(id int) (*models.Buku, error) {
	return bs.repo.GetBukuById(id)
}

func (bs *BukuService) GetAllBuku() ([]models.Buku, error) {
	return bs.repo.GetAllBuku()
}

func (bs *BukuService) AddBuku(buku models.Buku) error {
	return bs.repo.AddBuku(buku)
}

func (bs *BukuService) UpdateBuku(buku models.Buku) error {
	return bs.repo.UpdateBuku(buku)
}

func (bs *BukuService) DeleteBuku(id int) error {
	return bs.repo.DeleteBuku(id)
}
