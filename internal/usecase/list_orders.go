package usecase

import (
	"context"
	"github.com/marcosocram/fullcycle-clean-arch/internal/entity"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(repo entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{OrderRepository: repo}
}

func (uc *ListOrdersUseCase) Execute(ctx context.Context) ([]entity.Order, error) {
	return uc.OrderRepository.List(ctx)
}
