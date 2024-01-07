package response_test

import (
	"sagara/internal/delivery/response"
	"sagara/test/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderResponse(t *testing.T) {
	orderDTO := testdata.NewOrderDTO()
	order := testdata.NewOrder(orderDTO)

	res := response.MapOrderDomainToResponse(order)

	assert.Equal(t, order.Id.String(), res.Id)
	assert.Equal(t, order.PicName, res.PicName)
}
