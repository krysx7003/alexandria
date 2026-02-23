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

	series := models.Series{Name: "He Who Fights with Monsters", Book: 10, Ongoing: true}
	authors := []models.Author{
		{Name: "Shirtaloon"},
		{Name: "Travis Deverell"},
	}
	genres := []models.Genre{
		{Name: "LitRPG"},
		{Name: "Fantasy"},
		{Name: "Adventure"},
	}
	
    books := []models.Book{
		{ID: "1", Title: "Book 10", Series: series, Genres: genres, Authors: authors, Publisher: "Aethon Books"},
		{ID: "2", Title: "Book 11", Series: series, Genres: genres, Authors: authors, Publisher: "Aethon Books"},
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

	series := models.Series{Name: "He Who Fights with Monsters", Book: 10, Ongoing: true}
	authors := []models.Author{
		{Name: "Shirtaloon"},
		{Name: "Travis Deverell"},
	}
	genres := []models.Genre{
		{Name: "LitRPG"},
		{Name: "Fantasy"},
		{Name: "Adventure"},
	}
    
    // Fetch book by ID
    book := models.Book{
		ID: id,
		Title: "Book 10",
		Series: series,
		Genres: genres,
		Authors: authors,
		Publisher: "Aethon Books",
	}    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(book)
}
