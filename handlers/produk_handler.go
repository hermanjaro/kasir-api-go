package handlers

import (
	"encoding/json"
	"kasir-api/domain"
	"kasir-api/services"
	"net/http"
	"strconv"
	"strings"
)

type ProdukHandler struct {
	service services.ProdukService
}

func NewProdukHandler(service services.ProdukService) *ProdukHandler {
	return &ProdukHandler{service: service}
}

func (h *ProdukHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	produk, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produk)
}

func (h *ProdukHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	produk, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produk)
}

func (h *ProdukHandler) Create(w http.ResponseWriter, r *http.Request) {
	var produkBaru domain.Produk
	if err := json.NewDecoder(r.Body).Decode(&produkBaru); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	produk, err := h.service.Create(produkBaru)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(produk)
}

func (h *ProdukHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedItem domain.Produk
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	produk, err := h.service.Update(id, updatedItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produk)
}

func (h *ProdukHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Produk deleted successfully"})
}
