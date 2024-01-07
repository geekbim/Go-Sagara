package region_test

import (
	"context"
	"errors"
	"sagara/domain/entity"
	region_usecase "sagara/internal/usecase/region"
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

func TestListRegion(t *testing.T) {
	region := testdata.NewRegion()
	regions := []*entity.Region{region}

	err := errors.New("error")

	testCases := []struct {
		testName       string
		getRegion      funcCall
		countRegion    funcCall
		expectedResult []*entity.Region
		expectedError  []error
	}{
		{
			testName: "OK",
			getRegion: funcCall{
				Called: true,
				Input: []interface{}{
					mock.Anything, mock.Anything,
				},
				Output: []interface{}{
					regions, nil,
				},
			},
			countRegion: funcCall{
				Called: true,
				Input: []interface{}{
					mock.Anything, mock.Anything,
				},
				Output: []interface{}{
					int32(len(regions)), nil,
				},
			},
			expectedResult: regions,
		},
		{
			testName: "ErrorGetRegion",
			getRegion: funcCall{
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
			testName: "ErrorCountRegion",
			getRegion: funcCall{
				Called: true,
				Input: []interface{}{
					mock.Anything, mock.Anything,
				},
				Output: []interface{}{
					regions, nil,
				},
			},
			countRegion: funcCall{
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
			regionRepo := new(mocks.RegionRepository)

			if testCase.getRegion.Called {
				regionRepo.
					On("GetRegion", testCase.getRegion.Input...).
					Return(testCase.getRegion.Output...)
			}

			if testCase.countRegion.Called {
				regionRepo.
					On("CountRegion", testCase.countRegion.Input...).
					Return(testCase.countRegion.Output...)
			}

			useCase := region_usecase.NewRegionInteractor(regionRepo)

			res, total, err := useCase.ListRegion(context.Background(), nil)
			regionRepo.AssertExpectations(t)

			if err != nil {
				assert.Nil(t, res)
				assert.Equal(t, testCase.expectedError, err.Errors.Errors)
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, region.Name, res[0].Name)
			assert.Equal(t, int32(len(regions)), total)
		})
	}
}
