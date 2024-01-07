package usecase

import (
	"context"
	"sagara/domain/entity"
	"sagara/internal/delivery/request"
	"sagara/pkg/exceptions"
)

type RegionUseCase interface {
	ListRegion(ctx context.Context, options *request.Option) ([]*entity.Region, int32, *exceptions.CustomError)
}
