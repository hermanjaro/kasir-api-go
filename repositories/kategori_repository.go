package repositories

import (
	"kasir-api/config"
	"kasir-api/domain"
)

type KategoriRepository interface {
	GetAll() ([]domain.Kategori, error)
	GetByID(id int) (domain.Kategori, error)
	Create(kategori domain.Kategori) (domain.Kategori, error)
	Update(id int, kategori domain.Kategori) (domain.Kategori, error)
	Delete(id int) error
}

type kategoriRepository struct{}

func NewKategoriRepository() KategoriRepository {
	return &kategoriRepository{}
}

func (r *kategoriRepository) GetAll() ([]domain.Kategori, error) {
	query := `SELECT id, nama, deskripsi FROM kategori`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kategori []domain.Kategori
	for rows.Next() {
		var k domain.Kategori
		if err := rows.Scan(&k.ID, &k.Nama, &k.Deskripsi); err != nil {
			return nil, err
		}
		kategori = append(kategori, k)
	}

	return kategori, nil
}

func (r *kategoriRepository) GetByID(id int) (domain.Kategori, error) {
	query := `SELECT id, nama, deskripsi FROM kategori WHERE id = $1`
	var k domain.Kategori
	err := config.DB.QueryRow(query, id).Scan(&k.ID, &k.Nama, &k.Deskripsi)
	if err != nil {
		return domain.Kategori{}, err
	}
	return k, nil
}

func (r *kategoriRepository) Create(kategori domain.Kategori) (domain.Kategori, error) {
	query := `INSERT INTO kategori (nama, deskripsi) VALUES ($1, $2) RETURNING id`
	err := config.DB.QueryRow(query, kategori.Nama, kategori.Deskripsi).Scan(&kategori.ID)
	if err != nil {
		return domain.Kategori{}, err
	}
	return kategori, nil
}

func (r *kategoriRepository) Update(id int, kategori domain.Kategori) (domain.Kategori, error) {
	query := `UPDATE kategori SET nama = $1, deskripsi = $2 WHERE id = $3 RETURNING id, nama, deskripsi`
	err := config.DB.QueryRow(query, kategori.Nama, kategori.Deskripsi, id).
		Scan(&kategori.ID, &kategori.Nama, &kategori.Deskripsi)
	if err != nil {
		return domain.Kategori{}, err
	}
	return kategori, nil
}

func (r *kategoriRepository) Delete(id int) error {
	query := `DELETE FROM kategori WHERE id = $1`
	result, err := config.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return &NotFoundError{Message: "Kategori not found"}
	}
	return nil
}
