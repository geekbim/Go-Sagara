package user_test

import (
	"context"
	"errors"
	"sagara/domain/entity"
	user_usecase "sagara/internal/usecase/user"
	"sagara/mocks"
	"sagara/test/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	userDTO := testdata.NewUserDTO()
	user := testdata.NewUser(userDTO)
	user.Password = "qweasd123"

	err := errors.New("error")

	testCases := []struct {
		testName       string
		storeUser      funcCall
		expectedResult *entity.User
		expectedError  []error
	}{
		{
			testName: "OK",
			storeUser: funcCall{
				Called: true,
				Input: []interface{}{
					mock.Anything, mock.Anything,
				},
				Output: []interface{}{
					nil,
				},
			},
			expectedResult: user,
		},
		{
			testName: "ErrorStoreUser",
			storeUser: funcCall{
				Called: true,
				Input: []interface{}{
					mock.Anything, mock.Anything,
				},
				Output: []interface{}{
					err,
				},
			},
			expectedError: []error{
				err,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			userRepo := new(mocks.UserRepository)

			if testCase.storeUser.Called {
				userRepo.
					On("StoreUser", testCase.storeUser.Input...).
					Return(testCase.storeUser.Output...).
					Once()
			}

			useCase := user_usecase.NewUserInteractor(userRepo)

			res, err := useCase.Register(context.Background(), user)
			userRepo.AssertExpectations(t)

			if err != nil {
				assert.Nil(t, res)
				assert.Equal(t, testCase.expectedError, err.Errors.Errors)
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, user.Email, res.Email)
		})
	}
}
