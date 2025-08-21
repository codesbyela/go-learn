package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "gorm.io/gorm"
    "strconv"
    "api/models"
)

type handler struct {
    DB *gorm.DB
}

func New(db *gorm.DB) handler {
    return handler{db}
}

func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
    // Pagination parameters
    pageStr := r.URL.Query().Get("page")
    limitStr := r.URL.Query().Get("limit")
    offsetStr := r.URL.Query().Get("offset")
    page := 1
    limit := 10
    offset := 0
    if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
        limit = l
    }
    if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
        offset = o
    } else {
        if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
            page = p
        }
        offset = (page - 1) * limit
    }

    var books []models.Book
    var total int64
    h.DB.Model(&models.Book{}).Count(&total)
    h.DB.Limit(limit).Offset(offset).Find(&books)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "List of all books",
        "books":   books,
        "pagination": map[string]interface{}{
            "page":   page,
            "limit":  limit,
            "offset": offset,
            "total":  total,
            "pages":  (total + int64(limit) - 1) / int64(limit),
        },
    })
}

func (h handler) GetBook(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    var book models.Book
    h.DB.First(&book, id)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Book details",
        "book":    book,
    })
}

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    json.NewDecoder(r.Body).Decode(&book)
    h.DB.Create(&book)
    w.WriteHeader(http.StatusCreated)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Book added successfully",
        "book":    book,
    })
}

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    var book models.Book
    h.DB.First(&book, id)
    json.NewDecoder(r.Body).Decode(&book)
    h.DB.Save(&book)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Book updated successfully",
        "book":    book,
    })
}

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    var book models.Book
    result := h.DB.Delete(&book, id)
    if result.Error != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "error":   "Error deleting book",
            "status":  http.StatusInternalServerError,
            "message": "An error occurred while deleting the book.",
        })
        return
    }
    if result.RowsAffected == 0 {
        w.WriteHeader(http.StatusNotFound)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "error":      "Book not found",
            "status":     http.StatusNotFound,
            "message":    "No book found with the given ID",
            "code":       "BOOK_NOT_FOUND",
            "details":    "The book with the specified ID does not exist in the database.",
            "suggestion": "Please check the ID and try again.",
        })
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Book deleted successfully",
    })
}
