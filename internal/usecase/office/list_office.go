package office

import (
	"context"
	"sagara/domain/entity"
	"sagara/internal/delivery/request"
	"sagara/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
)

func (interactor *officeInteractor) ListOffice(ctx context.Context, options *request.Option) ([]*entity.Office, int32, *exceptions.CustomError) {
	var (
		offices  []*entity.Office
		total    int32
		errRepo  error
		multierr *multierror.Error
	)

	offices, errRepo = interactor.officeRepository.GetOffice(ctx, options)
	if errRepo != nil {
		multierr = multierror.Append(multierr, errRepo)
		return nil, int32(0), &exceptions.CustomError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	total, errRepo = interactor.officeRepository.CountOffice(ctx, options)
	if errRepo != nil {
		multierr = multierror.Append(multierr, errRepo)
		return nil, int32(0), &exceptions.CustomError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	return offices, total, nil
}
