package models

import "errors"

// Category structure
type Category struct {
	ID             string `json:"-"`
	OrganizationID string `json:"organization_id"`
	Name           string `json:"name,omitempty"`
	ListOrder      int64  `json:"list_order,omitempty"`
	// CreatedAt      time.Time     `json:"-"`
	// UpdatedAt      time.Time     `json:"-"`
	// DeletedAt      time.Time     `json:"-"`
}

// Categories array representation of Category
type Categories []Category

var categorySlice = Categories{}

func init() {

	var newCategory = Category{
		ID:             "1",
		OrganizationID: "23",
		Name:           "Kahvaltı",
		ListOrder:      5,
	}
	categorySlice = append(categorySlice, newCategory)
}

// GetMyCategory gets a category
func GetMyCategory(id string) (*Category, error) {
	for _, cat := range categorySlice {
		if id == cat.ID {
			return &cat, nil
		}
	}

	return nil, errors.New("not found")
}
