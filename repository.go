package main

import "errors"

// ProductRepository is an in-memory product database.
type ProductRepository struct {
	Products map[string]*Product
}

// ByID finds a product by its ID
func (u *ProductRepository) ByID(ID string) (*Product, error) {
	for _, product := range u.Products {
		if product.ID == ID {
			return product, nil
		}
	}
	return nil, ErrNotFound
}

var (
	// ErrNotFound represents a resource not found (404)
	ErrNotFound = errors.New("not found")

	// ErrUnauthorized represents a Forbidden (403)
	ErrUnauthorized = errors.New("unauthorized")

	// ErrEmpty is returned when input string is empty
	ErrEmpty = errors.New("empty string")
)
