package models

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Choice struct {
	ID       	   string        `json:"id"`
	CategoryID     string        `json:"category_id"`
	ProductID      string        `json:"product_id"`
	OptionID       string        `json:"option_id"`
	Name           string        `json:"name,omitempty"`
	ListOrder      int64         `json:"list_order,omitempty"`
	Price          float64       `json:"price,omitempty"`
}
