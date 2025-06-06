package application

import (
	"api/internal/domain"
	"api/internal/ports/outbound"
)

type GetItemUseCase struct {
	Repo outbound.ItemRepository
}

func NewGetItemUseCase(repo outbound.ItemRepository) *GetItemUseCase {
	return &GetItemUseCase{Repo: repo}
}

func (uc *GetItemUseCase) Execute(id string) (*domain.Item, error) {
	return uc.Repo.GetByID(id)
}
