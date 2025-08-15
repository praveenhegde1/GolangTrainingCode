package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// Define a struct to represent the resource
type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price float64 `json:"price"`
}

// Simulate a database with a map
var (
	items   = make(map[int]*Item)
	nextID  = 1
	mu      sync.Mutex // Mutex to handle concurrent access
)

// Create Item (POST /items)
func createItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newItem Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mu.Lock()
	newItem.ID = nextID
	items[nextID] = &newItem
	nextID++
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

// Read Item (GET /items/{id})
func getItemHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	item, exists := items[id]
	mu.Unlock()

	if !exists {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(item)
}

// Update Item (PUT /items/{id})
func updateItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	item, exists := items[id]
	mu.Unlock()

	if !exists {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	var updatedItem Item
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mu.Lock()
	item.Name = updatedItem.Name
	item.Price = updatedItem.Price
	mu.Unlock()

	json.NewEncoder(w).Encode(item)
}

// Delete Item (DELETE /items/{id})
func deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	_, exists := items[id]
	if exists {
		delete(items, id)
	}
	mu.Unlock()

	if !exists {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// List All Items (GET /items)
func listItemsHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	itemList := make([]*Item, 0, len(items))
	for _, item := range items {
		itemList = append(itemList, item)
	}

	json.NewEncoder(w).Encode(itemList)
}

func main() {
	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if r.URL.Query().Get("id") != "" {
				getItemHandler(w, r)
			} else {
				listItemsHandler(w, r)
			}
		case http.MethodPost:
			createItemHandler(w, r)
		case http.MethodPut:
			updateItemHandler(w, r)
		case http.MethodDelete:
			deleteItemHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
