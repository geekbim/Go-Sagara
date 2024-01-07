package repository

import (
	"context"
	"sagara/domain/entity"
	"sagara/internal/delivery/request"
)

type RegionRepository interface {
	GetRegion(ctx context.Context, options *request.Option) ([]*entity.Region, error)
	CountRegion(ctx context.Context, options *request.Option) (int32, error)
}
