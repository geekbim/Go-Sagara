package response_test

import (
	"sagara/internal/delivery/response"
	"sagara/test/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserResponse(t *testing.T) {
	userDTO := testdata.NewUserDTO()
	user := testdata.NewUser(userDTO)

	res := response.MapUserDomainToResponse(user, "token")

	assert.Equal(t, user.Id.String(), res.Id)
	assert.Equal(t, user.Email, res.Email)
	assert.Equal(t, user.Name, res.Name)
}
