package categories

import (
	"ecom_go_api/internal/json"
	"ecom_go_api/internal/models"
	"log"
	"net/http"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.ListCategories(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Write(w, http.StatusOK, categories)
}

func (h *handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.Read(r, &category); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateCategory(r.Context(), &category); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusCreated, category)
}
