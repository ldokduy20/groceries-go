package models

type ShoppingItem struct {
	ID       uint64  `json:"id"`
	ImageURL string  `json:"imageURL"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
}
