package handlers

import (
	"encoding/json"
	"kasir-api/domain"
	"kasir-api/services"
	"net/http"
	"strconv"
	"strings"
)

type KategoriHandler struct {
	service services.KategoriService
}

func NewKategoriHandler(service services.KategoriService) *KategoriHandler {
	return &KategoriHandler{service: service}
}

func (h *KategoriHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	kategori, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kategori)
}

func (h *KategoriHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/kategori/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	kategori, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kategori)
}

func (h *KategoriHandler) Create(w http.ResponseWriter, r *http.Request) {
	var kategoriBaru domain.Kategori
	if err := json.NewDecoder(r.Body).Decode(&kategoriBaru); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	kategori, err := h.service.Create(kategoriBaru)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(kategori)
}

func (h *KategoriHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/kategori/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedItem domain.Kategori
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	kategori, err := h.service.Update(id, updatedItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kategori)
}

func (h *KategoriHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/kategori/")
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
	json.NewEncoder(w).Encode(map[string]string{"message": "Kategori deleted successfully"})
}
