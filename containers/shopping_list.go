package containers

import (
	"errors"

	"github.com/ldokduy20/groceries-go/models"
)

type ShoppingList struct {
	Items  []models.ShoppingItem
	LastID uint64
}

func NewShoppingList() ShoppingList {
	return ShoppingList{
		Items:  make([]models.ShoppingItem, 0),
		LastID: 0,
	}
}

func (s *ShoppingList) AddShoppingItem(name string, price float32, image_url string) error {
	if price < 0 {
		return errors.New("price should not be negative")
	}
	if name == "" {
		return errors.New("name should not be empty")
	}
	item := models.ShoppingItem{
		Name:     name,
		Price:    price,
		ID:       s.LastID + 1,
		ImageURL: image_url,
	}
	s.LastID += 1
	s.Items = append(s.Items, item)
	return nil
}
