package entity

import (
	"errors"
	"sagara/pkg/common"
	"time"

	"github.com/hashicorp/go-multierror"
)

type Order struct {
	Id               common.ID
	PicName          string
	DateStart        time.Time
	DateEnd          time.Time
	Payment          string
	CapacityEmployee int
	Requirement      string
	Other            string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type OrderDTO struct {
	Id               *common.ID
	PicName          string
	DateStart        time.Time
	DateEnd          time.Time
	Payment          string
	CapacityEmployee int
	Requirement      string
	Other            string
}

func NewOrder(orderDTO *OrderDTO) (*Order, *multierror.Error) {
	var multierr *multierror.Error

	if orderDTO.Id == nil {
		id := common.NewID()
		orderDTO.Id = &id
	}

	order := &Order{
		Id:               *orderDTO.Id,
		PicName:          orderDTO.PicName,
		DateStart:        orderDTO.DateStart,
		DateEnd:          orderDTO.DateEnd,
		Payment:          orderDTO.Payment,
		CapacityEmployee: orderDTO.CapacityEmployee,
		Requirement:      orderDTO.Requirement,
		Other:            orderDTO.Other,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	if errValidate := order.Validate(); errValidate != nil {
		multierr = multierror.Append(multierr, errValidate)
	}

	if multierr != nil {
		return nil, multierr
	}

	return order, nil
}

func (order *Order) Validate() *multierror.Error {
	var multierr *multierror.Error

	if order.PicName == "" {
		multierr = multierror.Append(multierr, errors.New("pic name cannot be empty"))
	}

	return multierr
}
