package store

import (
	"github.com/shopping-list-backend/internal/models"
	"testing"
)

func TestItemStore_CRUD(t *testing.T) {
	store := NewItemStore()

	// Create
	item, _ := models.NewItem("Milk", "Dairy", "john", 2)
	err := store.Create(item)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	if item.ID == "" {
		t.Error("ID should be generated")
	}

	// Get
	retrieved, err := store.Get(item.ID)
	if err != nil || retrieved.Name != "Milk" {
		t.Error("Get failed")
	}

	// Update
	item.Name = "Almond Milk"
	err = store.Update(item.ID, item)
	if err != nil {
		t.Error("Update failed")
	}

	// Delete
	err = store.Delete(item.ID)
	if err != nil {
		t.Error("Delete failed")
	}

	// Should not exist
	_, err = store.Get(item.ID)
	if err != ErrItemNotFound {
		t.Error("Should return not found")
	}
}

func TestItemStore_GetAll(t *testing.T) {
	store := NewItemStore()

	item1, _ := models.NewItem("Milk", "Dairy", "john", 2)
	item2, _ := models.NewItem("Bread", "Bakery", "jane", 1)

	store.Create(item1)
	store.Create(item2)

	items := store.GetAll()
	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}
}
