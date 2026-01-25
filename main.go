package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// --- STRUCTS - Semacan tipe data gitu---

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

// --- DATA, menentukan data berdasarkan struct yang tersedia ---

var produk = []Produk{
	{ID: 1, Nama: "Wamena", Harga: 10000, Stok: 10},
	{ID: 2, Nama: "Toraja", Harga: 15000, Stok: 5},
	{ID: 3, Nama: "Enrekang", Harga: 20000, Stok: 8},
}

var kategori = []Kategori{
	{ID: 1, Nama: "Arabica", Deskripsi: "Kopi Arabica"},
	{ID: 2, Nama: "Robusta", Deskripsi: "Kopi jenis Robusta"},
	{ID: 3, Nama: "Blend", Deskripsi: "Campuran antara Arabica dan Robusta 50:50"},
}

// --- PRODUK HANDLERS ---

func getProdukByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, p := range produk {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.Error(w, "Produk tidak tersedia", http.StatusNotFound)
}

func updateProduk(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedItem Produk
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for i, p := range produk {
		if p.ID == id {
			produk[i].Nama = updatedItem.Nama
			produk[i].Harga = updatedItem.Harga
			produk[i].Stok = updatedItem.Stok

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk[i])
			return
		}
	}
	http.Error(w, "Produk not found", http.StatusNotFound)
}

func deleteProduk(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, p := range produk {
		if p.ID == id {
			produk = append(produk[:i], produk[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Produk deleted successfully"})
			return
		}
	}
	http.Error(w, "Produk not found", http.StatusNotFound)
}

// --- KATEGORI HANDLERS ---

func getKategoriByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/kategori/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, k := range kategori {
		if k.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(k)
			return
		}
	}
	http.Error(w, "Kategori tidak tersedia", http.StatusNotFound)
}

func updateKategori(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/kategori/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedItem Kategori
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for i, k := range kategori {
		if k.ID == id {
			kategori[i].Nama = updatedItem.Nama
			kategori[i].Deskripsi = updatedItem.Deskripsi

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(kategori[i])
			return
		}
	}
	http.Error(w, "Kategori not found", http.StatusNotFound)
}

func deleteKategori(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/kategori/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, k := range kategori {
		if k.ID == id {
			kategori = append(kategori[:i], kategori[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Kategori deleted successfully"})
			return
		}
	}
	http.Error(w, "Kategori not found", http.StatusNotFound)
}

// --- MAIN ---

func main() {
	// ================= PRODUK ROUTES =================

	// Handler untuk /api/produk/{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProdukByID(w, r)
		} else if r.Method == "PUT" {
			updateProduk(w, r)
		} else if r.Method == "DELETE" {
			deleteProduk(w, r)
		}
	})

	// Handler untuk /api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)
		} else if r.Method == "POST" {
			var produkBaru Produk
			if err := json.NewDecoder(r.Body).Decode(&produkBaru); err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}
			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(produkBaru)
		}
	})

	// ================= KATEGORI ROUTES =================

	// Handler untuk /api/kategori/{id}
	http.HandleFunc("/api/kategori/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getKategoriByID(w, r)
		} else if r.Method == "PUT" {
			updateKategori(w, r)
		} else if r.Method == "DELETE" {
			deleteKategori(w, r)
		}
	})

	// Handler untuk  /api/kategori
	http.HandleFunc("/api/kategori", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(kategori)
		} else if r.Method == "POST" {
			var kategoriBaru Kategori
			if err := json.NewDecoder(r.Body).Decode(&kategoriBaru); err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}
			kategoriBaru.ID = len(kategori) + 1
			kategori = append(kategori, kategoriBaru)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(kategoriBaru)
		}
	})

	// ================= SYSTEM ROUTES =================

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	fmt.Println("Server running di localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("gagal running server")
	}
}
