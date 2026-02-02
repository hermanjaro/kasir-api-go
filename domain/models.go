package domain

type Produk struct {
	ID            int    `json:"id"`
	Nama          string `json:"nama"`
	Harga         int    `json:"harga"`
	Stok          int    `json:"stok"`
	ID_Kategori    int    `json:"id_kategori"`
	Kategori_Nama string `json:"kategori_nama"`
}

type Kategori struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
}
