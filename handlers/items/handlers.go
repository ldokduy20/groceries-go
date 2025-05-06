package items

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ldokduy20/groceries-go/context"
	"github.com/ldokduy20/groceries-go/models"
)

type AddItemsResponse struct {
	ID      uint64 `json:"id"`
	Message string `json:"message"`
}

func ListItemsHandler(c *context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(c.ShoppingList.Items); err != nil {
			log.Printf("Error marshalling JSON: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}
}

func AddItemsHandler(c *context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var item models.ShoppingItem
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			log.Printf("Error unmarshalling JSON: %v", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
		}
		defer r.Body.Close()

		err := c.ShoppingList.AddShoppingItem(item.Name, item.Price, item.ImageURL)
		if err != nil {
			log.Printf("Shopping item is not valid: %v", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		response := AddItemsResponse{
			ID:      c.ShoppingList.LastID,
			Message: "Item created successfully!",
		}
		json.NewEncoder(w).Encode(response)
	}
}
