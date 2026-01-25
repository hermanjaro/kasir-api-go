package main

// --- Structs ---

type Produk struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stok  int    `json:"stok"`
}

type Kategori struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
}

// --- Data Store (In-Memory) ---

var produkList = []Produk{
	{ID: 1, Nama: "Wamena", Harga: 10000, Stok: 10},
	{ID: 2, Nama: "Toraja", Harga: 15000, Stok: 5},
	{ID: 3, Nama: "Enrekang", Harga: 20000, Stok: 8},
}

var kategoriList = []Kategori{
	{ID: 1, Nama: "Arabica", Deskripsi: "Kopi Arabica"},
	{ID: 2, Nama: "Robusta", Deskripsi: "Kopi jenis Robusta"},
	{ID: 3, Nama: "Blend", Deskripsi: "Campuran antara Arabica dan Robusta 50:50"},
}
