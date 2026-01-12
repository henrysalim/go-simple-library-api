package api

import (
	"encoding/json"
	"net/http"

	"simple-library-api/internal/model"
	"simple-library-api/internal/service"
)

type Handler struct {
	bookService *service.BookService
}

func NewHandler(bookService *service.BookService) *Handler {
	return &Handler{bookService: bookService}
}

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.bookService.CreateBook(&book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}
