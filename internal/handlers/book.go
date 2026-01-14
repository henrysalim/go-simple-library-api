package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"simple-library-api/internal/model"
	"simple-library-api/internal/repository"
	"strconv"
)

type BookHandler struct {
	Repo *repository.BookRepository
}

func NewBookHandler(repo *repository.BookRepository) *BookHandler {
	return &BookHandler{Repo: repo}
}

// CreateBook godoc
// @Summary      Create a new book
// @Description  Add a new book to the library
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book  body      model.Book  true  "Book JSON"
// @Success      201   {object}  model.Book
// @Failure      400   {string}  string "Invalid body"
// @Failure      500   {string}  string "Internal Server Error"
// @Router       /books [post]
func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := h.Repo.CreateBook(r.Context(), &book); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// GetBooks godoc
// @Summary      Get all books
// @Description  Get a list of all books stored in the database
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.Book
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /books [get]
func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	books, err := h.Repo.GetBooks(r.Context())
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// UpdateBook godoc
// @Summary      Update a book
// @Description  Update details of an existing book by ID
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "Book ID"
// @Param        book  body      model.Book  true  "Updated Book JSON"
// @Success      200   {object}  map[string]string
// @Failure      400   {string}  string "Invalid ID or Body"
// @Failure      404   {string}  string "Book not found"
// @Failure      500   {string}  string "Internal Server Error"
// @Router       /books/{id} [put]
func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid 'id' format", http.StatusBadRequest)
		return
	}

	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	book.ID = id

	if err := h.Repo.UpdateBook(r.Context(), &book); err != nil {
		if errors.Is(err, repository.ErrBookNotFound) {
			http.Error(w, "Book not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Book updated successfully!"})
}

// DeleteBook godoc
// @Summary      Delete a book
// @Description  Remove a book from the library by ID
// @Tags         books
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Success      200  {string}  string "Book deleted successfully"
// @Failure      400  {string}  string "Invalid ID"
// @Failure      404  {string}  string "Book not found"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /books/{id} [delete]
func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "Invalid 'id' format", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid 'id' format", http.StatusBadRequest)
		return
	}

	if err := h.Repo.DeleteBook(r.Context(), id); err != nil {
		if errors.Is(err, repository.ErrBookNotFound) {
			http.Error(w, "Book not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book deleted successfully!"))
}
