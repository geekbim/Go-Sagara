package usecase

import (
	"context"
	"sagara/domain/entity"
	"sagara/internal/delivery/request"
	"sagara/pkg/common"
	"sagara/pkg/exceptions"
)

type OfficeUseCase interface {
	ListOffice(ctx context.Context, options *request.Option) ([]*entity.Office, int32, *exceptions.CustomError)
	DetailOffice(ctx context.Context, id common.ID) (*entity.Office, *exceptions.CustomError)
}
