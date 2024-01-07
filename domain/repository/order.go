package repository

import (
	"context"
	"sagara/domain/entity"
)

type OrderRepository interface {
	StoreOrder(ctx context.Context, order *entity.Order) error
}
