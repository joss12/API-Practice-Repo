package models

import (
	"testing"
	"time"
)

func TestNewItem(t *testing.T) {
	tests := []struct {
		name        string
		itemName    string
		category    string
		createdBy   string
		quantity    int
		shouldError bool
		expectedErr error
	}{
		{
			name:        "valid item",
			itemName:    "Milk",
			category:    "Dairy",
			createdBy:   "john",
			quantity:    2,
			shouldError: false,
		},
		{
			name:        "empty name",
			itemName:    "",
			category:    "Dairy",
			createdBy:   "john",
			quantity:    1,
			shouldError: true,
			expectedErr: ErrEmptyName,
		},
		{
			name:        "whitespace name",
			itemName:    "   ",
			category:    "Dairy",
			createdBy:   "john",
			quantity:    1,
			shouldError: true,
			expectedErr: ErrEmptyName,
		},
		{
			name:        "zero quantity",
			itemName:    "Bread",
			category:    "Bakery",
			createdBy:   "john",
			quantity:    0,
			shouldError: true,
			expectedErr: ErrInvalidQuantity,
		},
		{
			name:        "negative quantity",
			itemName:    "Bread",
			category:    "Bakery",
			createdBy:   "john",
			quantity:    -5,
			shouldError: true,
			expectedErr: ErrInvalidQuantity,
		},
		{
			name:        "empty category",
			itemName:    "Eggs",
			category:    "",
			createdBy:   "john",
			quantity:    12,
			shouldError: true,
			expectedErr: ErrEmptyCategory,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item, err := NewItem(tt.itemName, tt.category, tt.createdBy, tt.quantity)

			if tt.shouldError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				if err != tt.expectedErr {
					t.Errorf("expected error %v, got %v", tt.expectedErr, err)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if item.Name != tt.itemName {
				t.Errorf("expected name %s, got %s", tt.itemName, item.Name)
			}

			if item.Quantity != tt.quantity {
				t.Errorf("expected quantity %d, got %d", tt.quantity, item.Quantity)
			}

			if item.Category != tt.category {
				t.Errorf("expected category %s, got %s", tt.category, item.Category)
			}

			if item.CreatedBy != tt.createdBy {
				t.Errorf("expected createdBy %s, got %s", tt.createdBy, item.CreatedBy)
			}

			if item.IsPurchased {
				t.Errorf("new item should not be purchased")
			}

			if item.CreatedAt.IsZero() {
				t.Errorf("CreatedAt should be set")
			}

			if item.UpdatedAt.IsZero() {
				t.Errorf("UpdatedAt should be set")
			}
		})
	}
}

func TestItem_MarkAsPurchased(t *testing.T) {
	item, _ := NewItem("Milk", "Dairy", "john", 2)

	if item.IsPurchased {
		t.Errorf("item should not be purchased initially")
	}

	oldUpdatedAt := item.UpdatedAt
	time.Sleep(10 * time.Millisecond) // Ensure time difference

	item.MarkAsPurchased()

	if !item.IsPurchased {
		t.Errorf("item should be marked as purchased")
	}

	if !item.UpdatedAt.After(oldUpdatedAt) {
		t.Errorf("UpdatedAt should be updated")
	}
}

func TestItem_MarkAsNotPurchased(t *testing.T) {
	item, _ := NewItem("Milk", "Dairy", "john", 2)
	item.MarkAsPurchased()

	if !item.IsPurchased {
		t.Errorf("item should be purchased")
	}

	oldUpdatedAt := item.UpdatedAt
	time.Sleep(10 * time.Millisecond)

	item.MarkAsNotPurchased()

	if item.IsPurchased {
		t.Errorf("item should not be marked as purchased")
	}

	if !item.UpdatedAt.After(oldUpdatedAt) {
		t.Errorf("UpdatedAt should be updated")
	}
}

func TestItem_UpdateQuantity(t *testing.T) {
	item, _ := NewItem("Milk", "Dairy", "john", 2)

	tests := []struct {
		name        string
		quantity    int
		shouldError bool
	}{
		{"valid update", 5, false},
		{"valid update to 1", 1, false},
		{"invalid zero", 0, true},
		{"invalid negative", -1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldUpdatedAt := item.UpdatedAt
			time.Sleep(10 * time.Millisecond)

			err := item.UpdateQuantity(tt.quantity)

			if tt.shouldError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if item.Quantity != tt.quantity {
				t.Errorf("expected quantity %d, got %d", tt.quantity, item.Quantity)
			}

			if !item.UpdatedAt.After(oldUpdatedAt) {
				t.Errorf("UpdatedAt should be updated")
			}
		})
	}
}

func TestItem_Validate(t *testing.T) {
	tests := []struct {
		name        string
		item        *Item
		shouldError bool
		expectedErr error
	}{
		{
			name: "valid item",
			item: &Item{
				Name:     "Milk",
				Quantity: 2,
				Category: "Dairy",
			},
			shouldError: false,
		},
		{
			name: "trims whitespace",
			item: &Item{
				Name:     "  Milk  ",
				Quantity: 2,
				Category: "  Dairy  ",
			},
			shouldError: false,
		},
		{
			name: "empty name after trim",
			item: &Item{
				Name:     "   ",
				Quantity: 2,
				Category: "Dairy",
			},
			shouldError: true,
			expectedErr: ErrEmptyName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.item.Validate()

			if tt.shouldError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				if err != tt.expectedErr {
					t.Errorf("expected error %v, got %v", tt.expectedErr, err)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
