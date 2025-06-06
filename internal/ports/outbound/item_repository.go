package outbound

import "api/internal/domain"

type ItemRepository interface {
	Save(item *domain.Item) error
	GetByID(id string) (*domain.Item, error)
}
