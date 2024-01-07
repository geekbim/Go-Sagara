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

type funcCall struct {
	Called bool
	Input  []interface{}
	Output []interface{}
}

func TestLogin(t *testing.T) {
	userDTO := testdata.NewUserDTO()
	user := testdata.NewUser(userDTO)
	user.Password = "qweasd123"
	user1 := testdata.NewUser(userDTO)
	user2 := testdata.NewUser(userDTO)
	user2.Password = "qweasd1234"

	err := errors.New("error")

	testCases := []struct {
		testName        string
		findUserByEmail funcCall
		expectedResult  *entity.User
		expectedError   []error
	}{
		{
			testName: "OK",
			findUserByEmail: funcCall{
				Called: true,
				Input: []interface{}{
					mock.Anything, mock.Anything,
				},
				Output: []interface{}{
					user1, nil,
				},
			},
			expectedResult: user,
		},
		{
			testName: "ErrorFindUser",
			findUserByEmail: funcCall{
				Called: true,
				Input: []interface{}{
					mock.Anything, mock.Anything,
				},
				Output: []interface{}{
					user1, err,
				},
			},
			expectedError: []error{
				err,
			},
		},
		{
			testName: "ErrorPasswordNotMatch",
			findUserByEmail: funcCall{
				Called: true,
				Input: []interface{}{
					mock.Anything, mock.Anything,
				},
				Output: []interface{}{
					user2, nil,
				},
			},
			expectedError: []error{
				errors.New("password does not match"),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			userRepo := new(mocks.UserRepository)

			if testCase.findUserByEmail.Called {
				userRepo.
					On("FindUserByEmail", testCase.findUserByEmail.Input...).
					Return(testCase.findUserByEmail.Output...)
			}

			useCase := user_usecase.NewUserInteractor(userRepo)

			res, err := useCase.Login(context.Background(), user)
			userRepo.AssertExpectations(t)

			if err != nil {
				assert.Nil(t, res)
				assert.Equal(t, testCase.expectedError, err.Errors.Errors)
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, user.Email, res.Email)
			assert.Equal(t, user1.Password, res.Password)
		})
	}
}
