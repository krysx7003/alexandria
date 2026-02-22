package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"alexandria/internal/models"
)

type BookHandler struct {
    // Could add database connection here
    // db *sql.DB
}

func NewBookHandler() *BookHandler {
    return &BookHandler{}
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
    // In real app, fetch from database
    books := []models.Book{
        {ID: "1", Title: "Book 1", Author: "Author 1", Year: "2023"},
        {ID: "2", Title: "Book 2", Author: "Author 2", Year: "2024"},
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    
    // Fetch book by ID
    book := models.Book{
        ID:     id,
        Title:  "Chuj",
        Author: "Test",
        Year:   "2137",
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(book)
}
