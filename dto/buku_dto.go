package dto

import "latian-clean-architecture/models"

type BukuDto struct {
	Id    *int   `json:"id" `
	Judul string `json:"judul"  `
	Harga *int   `json:"harga"  `
}

// ConvertToProductResponse mengubah model.Product menjadi dto.CreateProductRequest
func ConvertToBukuResponse(buku *models.Buku) BukuDto {
	return BukuDto{
		Id:    buku.Id,
		Judul: buku.Judul,
		Harga: buku.Harga,
	}
}
