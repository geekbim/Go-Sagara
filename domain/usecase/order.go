package usecase

import (
	"context"
	"sagara/domain/entity"
	"sagara/pkg/exceptions"
)

type OrderUseCase interface {
	CreateOrder(ctx context.Context, order *entity.Order) (*entity.Order, *exceptions.CustomError)
}
