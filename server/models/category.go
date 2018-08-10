package models

import (
	"errors"
	"strconv"
)

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
}

// ListCategory gets a category
func ListCategories() (*Categories, error) {
	return &CategorySlice, nil
}

// CreateCategory gets a category
func (category *Category) Create() (*Category, error) {
	category.ID = strconv.Itoa(len(CategorySlice) + 1)
	CategorySlice = append(CategorySlice, *category)

	return category, nil
}

// DeleteCategory gets a category
// func (category *Category) Delete(id string) (*Category, error) {
// 	for _, del := range CategorySlice {
// 		if id == del.ID {
// 			CategorySlice = delete(&Category, id)
// 			return &del, nil
// 		}
// 	}
// 	return category, nil
// }
