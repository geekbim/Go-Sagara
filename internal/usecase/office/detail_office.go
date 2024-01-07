package office

import (
	"context"
	"sagara/domain/entity"
	"sagara/pkg/common"
	"sagara/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
)

func (interactor *officeInteractor) DetailOffice(ctx context.Context, id common.ID) (*entity.Office, *exceptions.CustomError) {
	var (
		office   *entity.Office
		errRepo  error
		multierr *multierror.Error
	)

	office, errRepo = interactor.officeRepository.FindOfficeById(ctx, id)
	if errRepo != nil {
		multierr = multierror.Append(multierr, errRepo)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	return office, nil
}
