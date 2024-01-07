package entity

import (
	"errors"
	"sagara/pkg/common"
	"time"

	"github.com/hashicorp/go-multierror"
)

type User struct {
	Id             common.ID
	Email          string
	Name           string
	Password       string
	Position       string
	PhoneNumber    string
	CompanyName    string
	CompanyAddress string
	Information    string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type UserDTO struct {
	Id             *common.ID
	Name           string
	Email          string
	Password       string
	Position       string
	PhoneNumber    string
	CompanyName    string
	CompanyAddress string
	Information    string
}

func NewUser(userDTO *UserDTO) (*User, *multierror.Error) {
	var multierr *multierror.Error

	if userDTO.Id == nil {
		id := common.NewID()
		userDTO.Id = &id
	}

	user := &User{
		Id:             *userDTO.Id,
		Name:           userDTO.Name,
		Email:          userDTO.Email,
		Password:       userDTO.Password,
		Position:       userDTO.Position,
		PhoneNumber:    userDTO.PhoneNumber,
		CompanyName:    userDTO.CompanyName,
		CompanyAddress: userDTO.CompanyAddress,
		Information:    userDTO.Information,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if errValidate := user.Validate(); errValidate != nil {
		multierr = multierror.Append(multierr, errValidate)
	}

	if multierr != nil {
		return nil, multierr
	}

	return user, nil
}

func (user *User) Validate() *multierror.Error {
	var multierr *multierror.Error

	if user.Email == "" {
		multierr = multierror.Append(multierr, errors.New("email cannot be empty"))
	}

	if user.Password == "" {
		multierr = multierror.Append(multierr, errors.New("password cannot be empty"))
	}

	return multierr
}
