package models

import (
	"errors"
	"strings"
	"time"
)

type Item struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	Category    string    `json:"category"`
	IsPurchased bool      `json:"is_purchased"`
	CreatedBy   string    `json:"Created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Common validation errors
var (
	ErrEmptyName       = errors.New("item nam cannot be empty")
	ErrInvalidQuantity = errors.New("quantity must be greater than 0")
	ErrEmptyCategory   = errors.New("category cannot be empty")
)

// Validate checks if the item has valid data
func (i *Item) Validate() error {
	//Trim whitespace
	i.Name = strings.TrimSpace(i.Name)
	i.Category = strings.TrimSpace(i.Category)

	if i.Name == "" {
		return ErrEmptyName
	}

	if i.Quantity <= 0 {
		return ErrInvalidQuantity
	}

	if i.Category == "" {
		return ErrEmptyCategory
	}

	return nil
}

// MarkAsPurchased marks the item as purchased and updates timestamp
func (i *Item) MarkAsPurchased() {
	i.IsPurchased = true
	i.UpdatedAt = time.Now()
}

// MarkAsNotPurchased marks the item as not purchased and updates timestamp
func (i *Item) MarkAsNotPurchased() {
	i.IsPurchased = false
	i.UpdatedAt = time.Now()
}

// UpdateQuantity updates the item quantity with validation
func (i *Item) UpdateQuantity(quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}
	i.Quantity = quantity
	i.UpdatedAt = time.Now()
	return nil
}

// NewItem creates a new item with validation
func NewItem(name, category, createdby string, quantity int) (*Item, error) {
	now := time.Now()
	item := &Item{
		Name:        name,
		Quantity:    quantity,
		Category:    category,
		IsPurchased: false,
		CreatedBy:   createdby,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := item.Validate(); err != nil {
		return nil, err
	}

	return item, nil
}
