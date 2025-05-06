package context

import "github.com/ldokduy20/groceries-go/containers"

type Context struct {
	ShoppingList containers.ShoppingList
}

func NewContext() *Context {
	return &Context{
		ShoppingList: containers.NewShoppingList(),
	}
}
