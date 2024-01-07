package response_test

import (
	"sagara/domain/entity"
	"sagara/internal/delivery/response"
	"sagara/test/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegionResponse(t *testing.T) {
	region := testdata.NewRegion()
	regions := []*entity.Region{region}

	res := response.MapRegionListDomainToResponse(regions, int32(len(regions)))

	assert.Equal(t, regions[0].Id, res.Regions[0].Id)
	assert.Equal(t, regions[0].Name, res.Regions[0].Name)
}
