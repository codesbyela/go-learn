package handlers

import (
	"encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "gorm.io/gorm"
    "strconv"
    "ordermanagement/models"
)

type handler struct {
    DB *gorm.DB
}

func New(db *gorm.DB) handler {
    return handler{db}
}

func (h handler) GetProducts(w http.ResponseWriter, r *http.Request) {
    var products []models.Product
    q := h.DB.Model(&models.Product{})
    // Pagination parameters
    pageStr := r.URL.Query().Get("page")
    limitStr := r.URL.Query().Get("limit")
    page := 1
    limit := 10
    if pageStr != "" && limitStr != "" {
        p, err1 := strconv.Atoi(pageStr)
        l, err2 := strconv.Atoi(limitStr)
		page = p
		limit = l
        if err1 == nil && err2 == nil && page > 0 && limit > 0 {
            offset := (page - 1) * limit
            q = q.Where("quantity > 0").Offset(offset).Limit(limit)
        }
    }
    if err := q.Find(&products).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(map[string]interface{}{
        "products": products,
		"message":  "List of all products",
		"meta": map[string]interface{}{
			"page":  page,
			"limit": limit,
		},
    })
}

// get product
func (h handler) GetProduct(w http.ResponseWriter, r *http.Request) {
    idStr := mux.Vars(r)["id"]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var product models.Product
    if err := h.DB.First(&product, id).Error; err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(product)
}

// add product
func (h handler) AddProduct(w http.ResponseWriter, r *http.Request) {
    var product models.Product
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := h.DB.Create(&product).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(product)
}

// update product
func (h handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
    idStr := mux.Vars(r)["id"]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var product models.Product
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    product.ID = id
    if err := h.DB.Save(&product).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(product)
}


// delete product
func (h handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
    idStr := mux.Vars(r)["id"]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := h.DB.Delete(&models.Product{}, id).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}