package user

import (
	"context"
	"sagara/domain/entity"
	"sagara/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
)

func (interactor *userInteractor) Register(ctx context.Context, user *entity.User) (*entity.User, *exceptions.CustomError) {
	var multierr *multierror.Error

	errRepo := interactor.userRepository.StoreUser(ctx, user)
	if errRepo != nil {
		multierr = multierror.Append(multierr, errRepo)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multierr,
		}
	}

	return user, nil
}
