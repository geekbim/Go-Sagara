package order_test

import (
	"context"
	"errors"
	"sagara/domain/entity"
	order_usecase "sagara/internal/usecase/order"
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

func TestCreateOrder(t *testing.T) {
	orderDTO := testdata.NewOrderDTO()
	order := testdata.NewOrder(orderDTO)

	err := errors.New("error")

	testCases := []struct {
		testName       string
		storeOrder     funcCall
		expectedResult *entity.Order
		expectedError  []error
	}{
		{
			testName: "OK",
			storeOrder: funcCall{
				Called: true,
				Input: []interface{}{
					mock.Anything, mock.Anything,
				},
				Output: []interface{}{
					nil,
				},
			},
			expectedResult: order,
		},
		{
			testName: "ErrorStoreOrder",
			storeOrder: funcCall{
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
			orderRepo := new(mocks.OrderRepository)

			if testCase.storeOrder.Called {
				orderRepo.
					On("StoreOrder", testCase.storeOrder.Input...).
					Return(testCase.storeOrder.Output...).
					Once()
			}

			useCase := order_usecase.NewOrderInteractor(orderRepo)

			res, err := useCase.CreateOrder(context.Background(), order)
			orderRepo.AssertExpectations(t)

			if err != nil {
				assert.Nil(t, res)
				assert.Equal(t, testCase.expectedError, err.Errors.Errors)
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, order.PicName, res.PicName)
		})
	}
}
