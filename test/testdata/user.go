package testdata

import (
	"sagara/domain/entity"
	"sagara/internal/repository/psql/models"
	"sagara/pkg/common"
	"time"
)

func NewUserDTO() *entity.UserDTO {
	return &entity.UserDTO{
		Name:           "test",
		Email:          "test@email.com",
		Password:       "qweasd123",
		Position:       "test",
		PhoneNumber:    "81234567890",
		CompanyName:    "test",
		CompanyAddress: "test",
		Information:    "test",
	}
}

func NewUser(userDTO *entity.UserDTO) *entity.User {
	id, _ := common.StringToID("35da70af-aa50-44dc-ae6b-060a0f9e6933")
	return &entity.User{
		Id:             id,
		Name:           "test",
		Email:          userDTO.Email,
		Password:       "$2a$04$RrVwWOU9AjhQ3sO1UAYkE.98pEZRnXffcl7CRWskvejdqXBuiBm2i",
		Position:       "test",
		PhoneNumber:    "81234567890",
		CompanyName:    "test",
		CompanyAddress: "test",
		Information:    "test",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}
}

func NewUserModel() *models.User {
	return &models.User{
		Id:             "35da70af-aa50-44dc-ae6b-060a0f9e6933",
		Email:          "test@email.com",
		Name:           "test",
		Password:       "$2a$04$RrVwWOU9AjhQ3sO1UAYkE.98pEZRnXffcl7CRWskvejdqXBuiBm2i",
		Position:       "test",
		PhoneNumber:    "81234567890",
		CompanyName:    "test",
		CompanyAddress: "test",
		Information:    "test",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}
}
