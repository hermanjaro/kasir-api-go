package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/config"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"log"
	"net/http"
)

// --- MAIN ---

func main() {
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer config.CloseDB()

	produkRepo := repositories.NewProdukRepository()
	kategoriRepo := repositories.NewKategoriRepository()

	produkService := services.NewProdukService(produkRepo)
	kategoriService := services.NewKategoriService(kategoriRepo)

	produkHandler := handlers.NewProdukHandler(produkService)
	kategoriHandler := handlers.NewKategoriHandler(kategoriService)

	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			produkHandler.GetByID(w, r)
		case http.MethodPut:
			produkHandler.Update(w, r)
		case http.MethodDelete:
			produkHandler.Delete(w, r)
		}
	})

	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			produkHandler.GetAll(w, r)
		case http.MethodPost:
			produkHandler.Create(w, r)
		}
	})

	http.HandleFunc("/api/kategori/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			kategoriHandler.GetByID(w, r)
		case http.MethodPut:
			kategoriHandler.Update(w, r)
		case http.MethodDelete:
			kategoriHandler.Delete(w, r)
		}
	})

	http.HandleFunc("/api/kategori", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			kategoriHandler.GetAll(w, r)
		case http.MethodPost:
			kategoriHandler.Create(w, r)
		}
	})

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
