package application

import (
	"api/internal/domain"
	"api/internal/ports/outbound"

	"github.com/google/uuid"
)

type CreateItemUseCase struct {
	Repo outbound.ItemRepository
}

func NewCreateItemUseCase(repo outbound.ItemRepository) *CreateItemUseCase {
	return &CreateItemUseCase{Repo: repo}
}

func (uc *CreateItemUseCase) Execute(name string, value float64) (*domain.Item, error) {
	item := &domain.Item{
		ID:    uuid.New().String(),
		Name:  name,
		Value: value,
	}
	err := uc.Repo.Save(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}
