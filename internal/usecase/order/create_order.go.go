package order

import (
	"context"
	"sagara/domain/entity"
	"sagara/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
)

func (interactor *orderInteractor) CreateOrder(ctx context.Context, order *entity.Order) (*entity.Order, *exceptions.CustomError) {
	var multierr *multierror.Error

	errRepo := interactor.orderRepository.StoreOrder(ctx, order)
	if errRepo != nil {
		multierr = multierror.Append(multierr, errRepo)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multierr,
		}
	}

	return order, nil
}
