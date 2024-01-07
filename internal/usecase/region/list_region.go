package region

import (
	"context"
	"sagara/domain/entity"
	"sagara/internal/delivery/request"
	"sagara/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
)

func (interactor *regionInteractor) ListRegion(ctx context.Context, options *request.Option) ([]*entity.Region, int32, *exceptions.CustomError) {
	var (
		regions  []*entity.Region
		total    int32
		errRepo  error
		multierr *multierror.Error
	)

	regions, errRepo = interactor.regionRepository.GetRegion(ctx, options)
	if errRepo != nil {
		multierr = multierror.Append(multierr, errRepo)
		return nil, int32(0), &exceptions.CustomError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	total, errRepo = interactor.regionRepository.CountRegion(ctx, options)
	if errRepo != nil {
		multierr = multierror.Append(multierr, errRepo)
		return nil, int32(0), &exceptions.CustomError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	return regions, total, nil
}
