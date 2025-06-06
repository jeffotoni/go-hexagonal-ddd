package postgres

import (
	"api/internal/domain"
	"api/internal/ports/outbound"
	"errors"
	"sync"
)

type InMemoryItemRepository struct {
	mu    sync.RWMutex
	items map[string]*domain.Item
}

func NewInMemoryItemRepository() outbound.ItemRepository {
	return &InMemoryItemRepository{
		items: make(map[string]*domain.Item),
	}
}

func (r *InMemoryItemRepository) Save(item *domain.Item) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[item.ID] = item
	return nil
}

func (r *InMemoryItemRepository) GetByID(id string) (*domain.Item, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	item, ok := r.items[id]
	if !ok {
		return nil, errors.New("item not found")
	}
	return item, nil
}
