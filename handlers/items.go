package handlers

import (
	"net/http"

	"github.com/ldokduy20/groceries-go/context"
	"github.com/ldokduy20/groceries-go/handlers/items"
)

func ItemHandler(c *context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			items.ListItemsHandler(c)(w, r)
		case http.MethodPost:
			items.AddItemsHandler(c)(w, r)
		}
	}
}
