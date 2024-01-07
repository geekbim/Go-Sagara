package mapper_test

import (
	"sagara/domain/entity"
	"sagara/internal/repository/psql/mapper"
	"sagara/internal/repository/psql/models"
	"sagara/test/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegionMapper(t *testing.T) {
	regionDomain := testdata.NewRegion()
	regionsDomain := []*entity.Region{regionDomain}
	regionModel := testdata.NewRegionModel(regionDomain)
	regionsModel := []*models.Region{regionModel}
	regionListModel := testdata.NewRegionListModel()
	regionsListModel := []*models.RegionList{regionListModel}

	t.Run("ToDomainListRegion", func(t *testing.T) {
		res := mapper.ToDomainListRegion(regionsListModel)

		assert.Equal(t, res[0].Id, regionsDomain[0].Id)
		assert.Equal(t, res[0].Name, regionsDomain[0].Name)
	})

	t.Run("ToModelListeRegion", func(t *testing.T) {
		res := mapper.ToModelListRegion(regionsDomain)

		assert.Equal(t, res, regionsModel)
	})
}
