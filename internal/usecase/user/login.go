package user

import (
	"context"
	"errors"
	"sagara/domain/entity"
	"sagara/pkg/exceptions"
	"sagara/pkg/utils"

	"github.com/hashicorp/go-multierror"
)

func (interactor *userInteractor) Login(ctx context.Context, user *entity.User) (*entity.User, *exceptions.CustomError) {
	var multierr *multierror.Error

	res, errRepo := interactor.userRepository.FindUserByEmail(ctx, user.Email)
	if errRepo != nil {
		multierr = multierror.Append(multierr, errRepo)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	if !utils.CheckPasswordHash(user.Password, res.Password) {
		multierr = multierror.Append(multierr, errors.New("password does not match"))
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	return res, nil
}
