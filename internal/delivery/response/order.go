package response

import (
	"sagara/domain/entity"
	"time"
)

type Order struct {
	Id               string    `json:"id"`
	PicName          string    `json:"pic_name"`
	DateStart        time.Time `json:"date_start"`
	DateEnd          time.Time `json:"date_end"`
	Payment          string    `json:"payment"`
	CapacityEmployee int       `json:"capacity_employee"`
	Requirement      string    `json:"requirement"`
	Other            string    `json:"other"`
}

func MapOrderDomainToResponse(order *entity.Order) *Order {
	return &Order{
		Id:               order.Id.String(),
		PicName:          order.PicName,
		DateStart:        order.DateStart,
		DateEnd:          order.DateEnd,
		Payment:          order.Payment,
		CapacityEmployee: order.CapacityEmployee,
		Requirement:      order.Requirement,
		Other:            order.Other,
	}
}
