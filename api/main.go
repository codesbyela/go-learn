package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"api/pkg/db"
	"api/pkg/handlers"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Helper to load users from JSON file
func loadUsers() ([]User, error) {
	fileName := "users.json"
	var users []User

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return users, nil // no file = empty list
	}

	fileData, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	if len(fileData) == 0 {
		return users, nil
	}

	err = json.Unmarshal(fileData, &users)
	return users, err
}

// Helper to save users to JSON file
func saveUsers(users []User) error {
	fileData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("users.json", fileData, 0644)
}

// Create user (POST /user)
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	users, err := loadUsers()
	if err != nil {
		http.Error(w, "Error loading users", http.StatusInternalServerError)
		return
	}

	// Generate ID
	maxID := 0
	for _, u := range users {
		if u.ID > maxID {
			maxID = u.ID
		}
	}
	newUser.ID = maxID + 1

	users = append(users, newUser)

	err = saveUsers(users)
	if err != nil {
		http.Error(w, "Error saving user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created successfully",
		"data":    newUser,
	})
}

// Get all users (GET /users)
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := loadUsers()
	if err != nil {
		http.Error(w, "Error loading users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Users retrieved successfully",
		"data":    users,
	})
}

// Get user by ID (GET /user/{id})
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	users, err := loadUsers()
	if err != nil {
		http.Error(w, "Error loading users", http.StatusInternalServerError)
		return
	}

	for _, u := range users {
		if u.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "User found",
				"data":    u,
			})
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

// Update user by ID (PUT /user/{id})
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var updatedUser User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	users, err := loadUsers()
	if err != nil {
		http.Error(w, "Error loading users", http.StatusInternalServerError)
		return
	}

	found := false
	for i, u := range users {
		if u.ID == id {
			users[i].Name = updatedUser.Name
			users[i].Email = updatedUser.Email
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	err = saveUsers(users)
	if err != nil {
		http.Error(w, "Error saving user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User updated successfully",
	})
}

// Delete user by ID (DELETE /user/{id})
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	users, err := loadUsers()
	if err != nil {
		http.Error(w, "Error loading users", http.StatusInternalServerError)
		return
	}

	newUsers := []User{}
	found := false
	for _, u := range users {
		if u.ID != id {
			newUsers = append(newUsers, u)
		} else {
			found = true
		}
	}

	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	err = saveUsers(newUsers)
	if err != nil {
		http.Error(w, "Error saving users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User deleted successfully",
	})
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Welcome to the API",
	})
}

func main() {
	DB:= db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/", welcomeHandler).Methods("GET")
	router.HandleFunc("/user", createUserHandler).Methods("POST")
	router.HandleFunc("/users", getUsersHandler).Methods("GET")
	router.HandleFunc("/user/{id}", getUserHandler).Methods("GET")
	router.HandleFunc("/user/{id}/update", updateUserHandler).Methods("PUT")
	router.HandleFunc("/user/{id}/delete", deleteUserHandler).Methods("DELETE")

	// postgres routes
	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
    router.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
    router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
    router.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
    router.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
