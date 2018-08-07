package models

import "errors"

// Category structure
type Category struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	ListOrder  int64    `json:"list_order"`
	ProductIDs []string `json:"product_ids"`
	// CreatedAt      time.Time     `json:"-"`
	// UpdatedAt      time.Time     `json:"-"`
	// DeletedAt      time.Time     `json:"-"`
}

type Categories []Category

// GetCategory gets a category
func GetCategory(id string) (*Category, error) {
	for _, cat := range CategorySlice {
		if id == cat.ID {
			return &cat, nil
		}
	}

	return nil, errors.New("not found")
	// return nil, nil
}
