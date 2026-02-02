package repositories

import (
	"kasir-api/config"
	"kasir-api/domain"
)

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

type ProdukRepository interface {
	GetAll() ([]domain.Produk, error)
	GetByID(id int) (domain.Produk, error)
	Create(produk domain.Produk) (domain.Produk, error)
	Update(id int, produk domain.Produk) (domain.Produk, error)
	Delete(id int) error
}

type produkRepository struct{}

func NewProdukRepository() ProdukRepository {
	return &produkRepository{}
}

func (r *produkRepository) GetAll() ([]domain.Produk, error) {
	query := `SELECT p.id, p.nama, p.harga, p.stok, p.id_kategori, k.nama as kategori_nama
	          FROM produk p
	          LEFT JOIN kategori k ON p.id_kategori = k.id`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produk []domain.Produk
	for rows.Next() {
		var p domain.Produk
		if err := rows.Scan(&p.ID, &p.Nama, &p.Harga, &p.Stok, &p.ID_Kategori, &p.Kategori_Nama); err != nil {
			return nil, err
		}
		produk = append(produk, p)
	}

	return produk, nil
}

func (r *produkRepository) GetByID(id int) (domain.Produk, error) {
	query := `SELECT p.id, p.nama, p.harga, p.stok, p.id_kategori, k.nama as kategori_nama
	          FROM produk p
	          LEFT JOIN kategori k ON p.id_kategori = k.id
	          WHERE p.id = $1`
	var p domain.Produk
	err := config.DB.QueryRow(query, id).Scan(&p.ID, &p.Nama, &p.Harga, &p.Stok, &p.ID_Kategori, &p.Kategori_Nama)
	if err != nil {
		return domain.Produk{}, err
	}
	return p, nil
}

func (r *produkRepository) Create(produk domain.Produk) (domain.Produk, error) {
	query := `INSERT INTO produk (nama, harga, stok, id_kategori) VALUES ($1, $2, $3, $4) RETURNING id`
	err := config.DB.QueryRow(query, produk.Nama, produk.Harga, produk.Stok, produk.ID_Kategori).Scan(&produk.ID)
	if err != nil {
		return domain.Produk{}, err
	}

	// Fetch the created product with category name
	return r.GetByID(produk.ID)
}

func (r *produkRepository) Update(id int, produk domain.Produk) (domain.Produk, error) {
	query := `UPDATE produk SET nama = $1, harga = $2, stok = $3, id_kategori = $4 WHERE id = $5`
	_, err := config.DB.Exec(query, produk.Nama, produk.Harga, produk.Stok, produk.ID_Kategori, id)
	if err != nil {
		return domain.Produk{}, err
	}

	// Fetch the updated product with category name
	return r.GetByID(id)
}

func (r *produkRepository) Delete(id int) error {
	query := `DELETE FROM produk WHERE id = $1`
	result, err := config.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return &NotFoundError{Message: "Produk not found"}
	}
	return nil
}
