package store

import (
	"errors"
	"sync"

	"github.com/shopping-list-backend/internal/models"
)

var (
	ErrItemNotFound = errors.New("item not found")
	ErrItemExists   = errors.New("item already exists")
)

// ItemStore manages shopping list items in memory
type ItemStore struct {
	mu     sync.RWMutex
	items  map[string]*models.Item
	nextID int
}

// NewItemStore creates a new item store
func NewItemStore() *ItemStore {
	return &ItemStore{
		items:  make(map[string]*models.Item),
		nextID: 1,
	}
}

// Create adds a new item
func (s *ItemStore) Create(item *models.Item) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Generate ID
	item.ID = s.generateID()
	s.items[item.ID] = item
	return nil
}

// Get retrieves an item by ID
func (s *ItemStore) Get(id string) (*models.Item, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	item, exists := s.items[id]
	if !exists {
		return nil, ErrItemNotFound
	}
	return item, nil
}

// GetAll returns all items
func (s *ItemStore) GetAll() []*models.Item {
	s.mu.RLock()
	defer s.mu.RUnlock()

	items := make([]*models.Item, 0, len(s.items))
	for _, item := range s.items {
		items = append(items, item)
	}
	return items
}

// Update modifies an existing item
func (s *ItemStore) Update(id string, updatedItem *models.Item) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.items[id]; !exists {
		return ErrItemNotFound
	}

	updatedItem.ID = id
	s.items[id] = updatedItem
	return nil
}

// Delete removes an item
func (s *ItemStore) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.items[id]; !exists {
		return ErrItemNotFound
	}

	delete(s.items, id)
	return nil
}

// generateID creates a simple numeric ID
func (s *ItemStore) generateID() string {
	id := s.nextID
	s.nextID++
	return string(rune(id + '0')) // Simple ID like "1", "2", "3"
}
