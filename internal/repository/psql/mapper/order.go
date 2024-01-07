package mapper

import (
	"sagara/domain/entity"
	"sagara/internal/repository/psql/models"
	"sagara/pkg/common"

	"github.com/rocketlaunchr/dbq/v2"
)

func ToDomainOrder(m *models.Order) *entity.Order {
	id, _ := common.StringToID(m.Id)
	order := &entity.Order{
		Id:               id,
		PicName:          m.PicName,
		DateStart:        m.DateStart,
		DateEnd:          m.DateEnd,
		Payment:          m.Payment,
		CapacityEmployee: m.CapacityEmployee,
		Requirement:      m.Requirement,
		Other:            m.Other,
		CreatedAt:        m.CreatedAt,
		UpdatedAt:        m.UpdatedAt,
	}

	return order
}

func ToModelOrder(d *entity.Order) *models.Order {
	order := &models.Order{
		Id:               d.Id.String(),
		PicName:          d.PicName,
		DateStart:        d.DateStart,
		DateEnd:          d.DateEnd,
		Payment:          d.Payment,
		CapacityEmployee: d.CapacityEmployee,
		Requirement:      d.Requirement,
		Other:            d.Other,
		CreatedAt:        d.CreatedAt,
		UpdatedAt:        d.UpdatedAt,
	}

	return order
}

func DataDbqOrder(domain *entity.Order) []interface{} {
	return dbq.Struct(ToModelOrder(domain))
}

func ToDbqStructOrder(domain *entity.Order) (dbqStruct []interface{}) {
	dbqStruct = append(dbqStruct, DataDbqOrder(domain))
	return
}
