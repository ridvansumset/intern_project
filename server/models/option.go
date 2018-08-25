package models

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Option struct {
	ID       	   string           `json:"id"`
	CategoryID     string           `json:"category_id"`
	ProductID      string           `json:"product_id"`
	Name           string           `json:"name,omitempty"`
	ListOrder      int64            `json:"list_order,omitempty"`
	Required       bool             `json:"required,omitempty"`
	ChoiceType     OptionChoiceType `json:"choice_type,omitempty"`
}

type OptionChoiceType string

const (
	// OptionChoiceTypeSingle should used for single choices
	OptionChoiceTypeSingle OptionChoiceType = "single"
	// OptionChoiceTypeMultiple should used for multiple choices
	OptionChoiceTypeMultiple OptionChoiceType = "multiple"
)
