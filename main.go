package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// --- Produk Routes ---
	http.HandleFunc("/api/produk", HandleProduk)        // List & Create
	http.HandleFunc("/api/produk/", HandleProdukDetail) // Get, Update, Delete (ID)

	// --- Kategori Routes ---
	http.HandleFunc("/api/kategori", HandleKategori)        // List & Create
	http.HandleFunc("/api/kategori/", HandleKategoriDetail) // Get, Update, Delete (ID)

	// --- System Routes ---
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
