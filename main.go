package main

import (
	"log"
	"net/http"

	"github.com/ldokduy20/groceries-go/context"
	"github.com/ldokduy20/groceries-go/handlers"
)

func main() {
	ctx := context.NewContext()
	http.HandleFunc("/items/", handlers.ItemHandler(ctx))

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
