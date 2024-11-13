package services

import (
	"latian-clean-architecture/dto"
	"latian-clean-architecture/models"
	"latian-clean-architecture/repositories"
)

// BukuServiceInterface mendefinisikan kontrak yang harus dipenuhi oleh BukuService.
type BukuServiceInterface interface {
	GetBukuById(id int) (*dto.BukuDto, error)
	GetAllBuku() ([]dto.BukuDto, error)
	AddBuku(buku models.Buku) error
	UpdateBuku(buku models.Buku) error
	DeleteBuku(id int) error
}
type BukuService struct {
	repo repositories.BukuRepository
}

func NewBukuService(repo repositories.BukuRepository) *BukuService {
	return &BukuService{repo: repo}
}

func (bs *BukuService) GetBukuById(id int) (*dto.BukuDto, error) {
	// Ambil data buku dari repository
	buku, err := bs.repo.GetBukuById(id)
	if err != nil {
		return nil, err
	}

	// Konversi data buku ke bentuk yang lebih sesuai untuk respons

	bukuResponses := dto.ConvertToBukuResponse(buku)

	return &bukuResponses, nil
}

func (bs *BukuService) GetAllBuku() ([]dto.BukuDto, error) {
	// Ambil data buku dari repository
	bukuList, err := bs.repo.GetAllBuku()
	if err != nil {
		return nil, err
	}

	// Konversi data buku ke bentuk yang lebih sesuai untuk respons
	var bukuResponses []dto.BukuDto
	for _, buku := range bukuList {
		bukuResponses = append(bukuResponses, dto.ConvertToBukuResponse(&buku))
	}

	return bukuResponses, nil
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
