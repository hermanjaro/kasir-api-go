package services

import (
	"kasir-api/domain"
	"kasir-api/repositories"
)

type ProdukService interface {
	GetAll() ([]domain.Produk, error)
	GetByID(id int) (domain.Produk, error)
	Create(produk domain.Produk) (domain.Produk, error)
	Update(id int, produk domain.Produk) (domain.Produk, error)
	Delete(id int) error
}

type produkService struct {
	repo repositories.ProdukRepository
}

func NewProdukService(repo repositories.ProdukRepository) ProdukService {
	return &produkService{repo: repo}
}

func (s *produkService) GetAll() ([]domain.Produk, error) {
	return s.repo.GetAll()
}

func (s *produkService) GetByID(id int) (domain.Produk, error) {
	produk, err := s.repo.GetByID(id)
	if err != nil {
		return domain.Produk{}, &repositories.NotFoundError{Message: "Produk tidak tersedia"}
	}
	return produk, nil
}

func (s *produkService) Create(produk domain.Produk) (domain.Produk, error) {
	return s.repo.Create(produk)
}

func (s *produkService) Update(id int, produk domain.Produk) (domain.Produk, error) {
	updatedProduk, err := s.repo.Update(id, produk)
	if err != nil {
		return domain.Produk{}, &repositories.NotFoundError{Message: "Produk not found"}
	}
	return updatedProduk, nil
}

func (s *produkService) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return &repositories.NotFoundError{Message: "Produk not found"}
	}
	return nil
}
