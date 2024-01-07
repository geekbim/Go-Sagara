package mapper_test

import (
	"sagara/internal/repository/psql/mapper"
	"sagara/test/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserMapper(t *testing.T) {
	userModel := testdata.NewUserModel()
	userDTO := testdata.NewUserDTO()
	userDomain := testdata.NewUser(userDTO)

	t.Run("ToDomainUser", func(t *testing.T) {
		res := mapper.ToDomainUser(userModel)

		assert.Equal(t, res, userDomain)
	})

	t.Run("ToModelUser", func(t *testing.T) {
		res := mapper.ToModelUser(userDomain)

		assert.Equal(t, res, userModel)
	})
}
