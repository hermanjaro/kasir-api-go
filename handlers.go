package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// ================= PRODUK HANDLERS =================

// HandleProduk handles GET (List) and POST (Create)
func HandleProduk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(produkList)

	case http.MethodPost:
		var produkBaru Produk
		if err := json.NewDecoder(r.Body).Decode(&produkBaru); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		produkBaru.ID = len(produkList) + 1
		produkList = append(produkList, produkBaru)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(produkBaru)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// HandleProdukDetail handles GET (One), PUT (Update), DELETE (Remove)
func HandleProdukDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract ID
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Find Index
	index := -1
	for i, p := range produkList {
		if p.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "Produk not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(produkList[index])

	case http.MethodPut:
		var updatedItem Produk
		if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		// Update fields
		produkList[index].Nama = updatedItem.Nama
		produkList[index].Harga = updatedItem.Harga
		produkList[index].Stok = updatedItem.Stok
		json.NewEncoder(w).Encode(produkList[index])

	case http.MethodDelete:
		// Remove from slice
		produkList = append(produkList[:index], produkList[index+1:]...)
		json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf("Produk ID %d deleted", id),
		})

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// ================= KATEGORI HANDLERS =================

// HandleKategori handles GET (List) and POST (Create)
func HandleKategori(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(kategoriList)

	case http.MethodPost:
		var kategoriBaru Kategori
		if err := json.NewDecoder(r.Body).Decode(&kategoriBaru); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		kategoriBaru.ID = len(kategoriList) + 1
		kategoriList = append(kategoriList, kategoriBaru)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(kategoriBaru)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// HandleKategoriDetail handles GET (One), PUT (Update), DELETE (Remove)
func HandleKategoriDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/kategori/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	index := -1
	for i, k := range kategoriList {
		if k.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "Kategori not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(kategoriList[index])

	case http.MethodPut:
		var updatedItem Kategori
		if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		kategoriList[index].Nama = updatedItem.Nama
		kategoriList[index].Deskripsi = updatedItem.Deskripsi
		json.NewEncoder(w).Encode(kategoriList[index])

	case http.MethodDelete:
		kategoriList = append(kategoriList[:index], kategoriList[index+1:]...)
		json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf("Kategori ID %d deleted", id),
		})

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
