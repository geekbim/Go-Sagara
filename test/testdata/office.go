package testdata

import (
	"sagara/domain/entity"
	"sagara/internal/repository/psql/mapper"
	"sagara/internal/repository/psql/models"
	"sagara/pkg/common"
	"time"
)

func NewOffice() *entity.Office {
	id, _ := common.StringToID("bf1796fe-cc7c-4a46-a422-da7317c8916f")
	image := &entity.DataObject{
		Id:        id,
		Path:      "test",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	return &entity.Office{
		Id:          id,
		Name:        "test",
		RoomType:    "test",
		Description: "test",
		Image:       image,
		Region: &entity.Region{
			Id:        1,
			Name:      "Jakarta",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
		AdditionalData: &entity.OfficeAdditionalData{
			Capacity: 1,
			Table:    1,
			Chair:    1,
			SpecialOffer: &entity.SpecialOfferAdditionalData{
				Title:       "test",
				Description: "test",
				Image:       image,
			},
			SpecialPrice: &entity.SpecialPriceAdditionalData{
				Image: image,
			},
		},
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func NewOfficeModel(office *entity.Office) *models.Office {
	return mapper.ToModelOffice(office)
}
