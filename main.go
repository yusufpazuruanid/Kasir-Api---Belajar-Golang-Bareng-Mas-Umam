package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Category model
type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// In-memory storage
var categories = []Category{
	{ID: 1, Name: "Elektronik", Description: "Perangkat elektronik dan gadget"},
	{ID: 2, Name: "Makanan", Description: "Makanan dan minuman ringan"},
}

var nextID = 3

// Handle GET /categories and POST /categories
func categoriesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		// Get all categories
		json.NewEncoder(w).Encode(categories)

	case http.MethodPost:
		// Create new category
		var newCat Category
		err := json.NewDecoder(r.Body).Decode(&newCat)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		newCat.ID = nextID
		nextID++
		categories = append(categories, newCat)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newCat)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
	}
}

// Handle GET, PUT, DELETE /categories/{id}
func categoryByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract ID from URL
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// Get category by ID
		for _, cat := range categories {
			if cat.ID == id {
				json.NewEncoder(w).Encode(cat)
				return
			}
		}
		http.Error(w, "Category not found", http.StatusNotFound)

	case http.MethodPut:
		// Update category by ID
		var updatedCat Category
		err := json.NewDecoder(r.Body).Decode(&updatedCat)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		for i, cat := range categories {
			if cat.ID == id {
				updatedCat.ID = id // Keep the same ID
				categories[i] = updatedCat
				json.NewEncoder(w).Encode(updatedCat)
				return
			}
		}
		http.Error(w, "Category not found", http.StatusNotFound)

	case http.MethodDelete:
		// Delete category by ID
		for i, cat := range categories {
			if cat.ID == id {
				// Remove from slice
				categories = append(categories[:i], categories[i+1:]...)
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, `{"message": "Category with ID %d deleted"}`, id)
				return
			}
		}
		http.Error(w, "Category not found", http.StatusNotFound)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
	}
}

// localhost:8080/health
func healthHandler(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	}
// localhost:8080
func app(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"learn":  "Code With Umam - Belajar Membuat API Dengan Golang",
			"task": "Kasir API - Categories",
			"completed": "Yes",
			"completed_at": "Kamis, 3 Sya'ban 1447 H / 22 Januari 2026 M At 09:47:13",
		})
	}

func main() {
	// Router
	http.HandleFunc("/", app)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/categories", categoriesHandler)
	http.HandleFunc("/categories/", categoryByIDHandler)

	// Server setup
	port := ":8080"
	fmt.Printf("Server running on http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

