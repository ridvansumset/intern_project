package models

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

// GetMyCategory gets a category
func GetMyCategory(id string) (*Category, error) {
	// for _, cat := range categorySlice {
	// 	if id == cat.ID {
	// 		return &cat, nil
	// 	}
	// }

	// return nil, errors.New("not found")
	return nil, nil
}
