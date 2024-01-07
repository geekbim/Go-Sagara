package office_test

import (
	"context"
	"errors"
	"sagara/domain/entity"
	office_usecase "sagara/internal/usecase/office"
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

func TestListOffice(t *testing.T) {
	office := testdata.NewOffice()
	offices := []*entity.Office{office}

	err := errors.New("error")

	testCases := []struct {
		testName       string
		getOffice      funcCall
		countOffice    funcCall
		expectedResult []*entity.Office
		expectedError  []error
	}{
		{
			testName: "OK",
			getOffice: funcCall{
				Called: true,
				Input: []interface{}{
					mock.Anything, mock.Anything,
				},
				Output: []interface{}{
					offices, nil,
				},
			},
			countOffice: funcCall{
				Called: true,
				Input: []interface{}{
					mock.Anything, mock.Anything,
				},
				Output: []interface{}{
					int32(len(offices)), nil,
				},
			},
			expectedResult: offices,
		},
		{
			testName: "ErrorGetOffice",
			getOffice: funcCall{
				Called: true,
				Input: []interface{}{
					mock.Anything, mock.Anything,
				},
				Output: []interface{}{
					nil, err,
				},
			},
			expectedError: []error{
				err,
			},
		},
		{
			testName: "ErrorCountOffice",
			getOffice: funcCall{
				Called: true,
				Input: []interface{}{
					mock.Anything, mock.Anything,
				},
				Output: []interface{}{
					offices, nil,
				},
			},
			countOffice: funcCall{
				Called: true,
				Input: []interface{}{
					mock.Anything, mock.Anything,
				},
				Output: []interface{}{
					int32(0), err,
				},
			},
			expectedError: []error{
				err,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			officeRepo := new(mocks.OfficeRepository)

			if testCase.getOffice.Called {
				officeRepo.
					On("GetOffice", testCase.getOffice.Input...).
					Return(testCase.getOffice.Output...)
			}

			if testCase.countOffice.Called {
				officeRepo.
					On("CountOffice", testCase.countOffice.Input...).
					Return(testCase.countOffice.Output...)
			}

			useCase := office_usecase.NewOfficeInteractor(officeRepo)

			res, total, err := useCase.ListOffice(context.Background(), nil)
			officeRepo.AssertExpectations(t)

			if err != nil {
				assert.Nil(t, res)
				assert.Equal(t, testCase.expectedError, err.Errors.Errors)
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, office.Name, res[0].Name)
			assert.Equal(t, int32(len(offices)), total)
		})
	}
}
