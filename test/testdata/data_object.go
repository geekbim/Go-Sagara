package testdata

import (
	"sagara/domain/entity"
	"sagara/internal/repository/psql/models"
	"sagara/pkg/common"
	"time"
)

func NewDataObject() *entity.DataObject {
	id, _ := common.StringToID("35da70af-aa50-44dc-ae6b-060a0f9e6933")
	return &entity.DataObject{
		Id:        id,
		Path:      "/static/office/serviced-office.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func NewDataObjectModel() *models.DataObject {
	return &models.DataObject{
		Id:        "35da70af-aa50-44dc-ae6b-060a0f9e6933",
		Path:      "/static/office/serviced-office.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}
