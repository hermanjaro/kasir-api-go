package services

import (
	"kasir-api/domain"
	"kasir-api/repositories"
)

type KategoriService interface {
	GetAll() ([]domain.Kategori, error)
	GetByID(id int) (domain.Kategori, error)
	Create(kategori domain.Kategori) (domain.Kategori, error)
	Update(id int, kategori domain.Kategori) (domain.Kategori, error)
	Delete(id int) error
}

type kategoriService struct {
	repo repositories.KategoriRepository
}

func NewKategoriService(repo repositories.KategoriRepository) KategoriService {
	return &kategoriService{repo: repo}
}

func (s *kategoriService) GetAll() ([]domain.Kategori, error) {
	return s.repo.GetAll()
}

func (s *kategoriService) GetByID(id int) (domain.Kategori, error) {
	kategori, err := s.repo.GetByID(id)
	if err != nil {
		return domain.Kategori{}, &repositories.NotFoundError{Message: "Kategori tidak tersedia"}
	}
	return kategori, nil
}

func (s *kategoriService) Create(kategori domain.Kategori) (domain.Kategori, error) {
	return s.repo.Create(kategori)
}

func (s *kategoriService) Update(id int, kategori domain.Kategori) (domain.Kategori, error) {
	updatedKategori, err := s.repo.Update(id, kategori)
	if err != nil {
		return domain.Kategori{}, &repositories.NotFoundError{Message: "Kategori not found"}
	}
	return updatedKategori, nil
}

func (s *kategoriService) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return &repositories.NotFoundError{Message: "Kategori not found"}
	}
	return nil
}
