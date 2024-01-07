package mapper_test

import (
	"sagara/internal/repository/psql/mapper"
	"sagara/test/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderMapper(t *testing.T) {
	orderModel := testdata.NewOrderModel()
	orderDTO := testdata.NewOrderDTO()
	orderDomain := testdata.NewOrder(orderDTO)

	t.Run("ToDomainOrder", func(t *testing.T) {
		res := mapper.ToDomainOrder(orderModel)

		assert.Equal(t, res, orderDomain)
	})

	t.Run("ToModelOrder", func(t *testing.T) {
		res := mapper.ToModelOrder(orderDomain)

		assert.Equal(t, res, orderModel)
	})
}
