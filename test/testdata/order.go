package testdata

import (
	"sagara/domain/entity"
	"sagara/internal/repository/psql/models"
	"sagara/pkg/common"
	"time"
)

func NewOrderDTO() *entity.OrderDTO {
	return &entity.OrderDTO{
		PicName:          "test",
		DateStart:        time.Time{},
		DateEnd:          time.Time{},
		Payment:          "test",
		CapacityEmployee: 1,
		Requirement:      "test",
		Other:            "test",
	}
}

func NewOrder(orderDTO *entity.OrderDTO) *entity.Order {
	id, _ := common.StringToID("35da70af-aa50-44dc-ae6b-060a0f9e6933")
	return &entity.Order{
		Id:               id,
		PicName:          "test",
		DateStart:        time.Time{},
		DateEnd:          time.Time{},
		Payment:          "test",
		CapacityEmployee: 1,
		Requirement:      "test",
		Other:            "test",
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
	}
}

func NewOrderModel() *models.Order {
	return &models.Order{
		Id:               "35da70af-aa50-44dc-ae6b-060a0f9e6933",
		PicName:          "test",
		DateStart:        time.Time{},
		DateEnd:          time.Time{},
		Payment:          "test",
		CapacityEmployee: 1,
		Requirement:      "test",
		Other:            "test",
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
	}
}
