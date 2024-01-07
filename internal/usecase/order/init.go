package order

import (
	"sagara/domain/repository"
	"sagara/domain/usecase"
)

type orderInteractor struct {
	orderRepository repository.OrderRepository
}

func NewOrderInteractor(
	orderRepository repository.OrderRepository,
) usecase.OrderUseCase {
	return &orderInteractor{
		orderRepository: orderRepository,
	}
}
