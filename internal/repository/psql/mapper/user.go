package mapper

import (
	"sagara/domain/entity"
	"sagara/internal/repository/psql/models"
	"sagara/pkg/common"

	"github.com/rocketlaunchr/dbq/v2"
)

func ToDomainUser(m *models.User) *entity.User {
	id, _ := common.StringToID(m.Id)
	user := &entity.User{
		Id:             id,
		Name:           m.Name,
		Email:          m.Email,
		Password:       m.Password,
		Position:       m.Position,
		PhoneNumber:    m.PhoneNumber,
		CompanyName:    m.CompanyName,
		CompanyAddress: m.CompanyAddress,
		Information:    m.Information,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
	}

	return user
}

func ToModelUser(d *entity.User) *models.User {
	user := &models.User{
		Id:             d.Id.String(),
		Name:           d.Name,
		Email:          d.Email,
		Password:       d.Password,
		Position:       d.Position,
		PhoneNumber:    d.PhoneNumber,
		CompanyName:    d.CompanyName,
		CompanyAddress: d.CompanyAddress,
		Information:    d.Information,
		CreatedAt:      d.CreatedAt,
		UpdatedAt:      d.UpdatedAt,
	}

	return user
}

func DataDbqUser(domain *entity.User) []interface{} {
	return dbq.Struct(ToModelUser(domain))
}

func ToDbqStructUser(domain *entity.User) (dbqStruct []interface{}) {
	dbqStruct = append(dbqStruct, DataDbqUser(domain))
	return
}
