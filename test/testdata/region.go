package testdata

import (
	"sagara/domain/entity"
	"sagara/internal/repository/psql/mapper"
	"sagara/internal/repository/psql/models"
	"time"
)

func NewRegion() *entity.Region {
	return &entity.Region{
		Id:        1,
		Name:      "Jakarta",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func NewRegionModel(region *entity.Region) *models.Region {
	return mapper.ToModelRegion(region)
}

func NewRegionListModel() *models.RegionList {
	return &models.RegionList{
		Id:        1,
		Name:      "Jakarta",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}
