package mapper_test

import (
	"sagara/internal/repository/psql/mapper"
	"sagara/test/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataObjectMapper(t *testing.T) {
	dataObjectModel := testdata.NewDataObjectModel()
	dataObjectDomain := testdata.NewDataObject()

	res := mapper.ToDomainDataObject(dataObjectModel)

	assert.Equal(t, res, dataObjectDomain)
}
