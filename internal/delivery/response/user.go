package response

import "sagara/domain/entity"

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func MapUserDomainToResponse(user *entity.User, token string) *User {
	return &User{
		Id:    user.Id.String(),
		Email: user.Email,
		Name:  user.Name,
		Token: token,
	}
}
