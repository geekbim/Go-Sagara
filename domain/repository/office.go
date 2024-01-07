package repository

import (
	"context"
	"sagara/domain/entity"
	"sagara/internal/delivery/request"
	"sagara/pkg/common"
)

type OfficeRepository interface {
	GetOffice(ctx context.Context, options *request.Option) ([]*entity.Office, error)
	CountOffice(ctx context.Context, options *request.Option) (int32, error)
	FindOfficeById(ctx context.Context, id common.ID) (*entity.Office, error)
}
