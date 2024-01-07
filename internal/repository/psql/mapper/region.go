package mapper

import (
	"sagara/domain/entity"
	"sagara/internal/repository/psql/models"

	"github.com/rocketlaunchr/dbq/v2"
)

func ToDomainRegion(m *models.RegionList) *entity.Region {
	region := &entity.Region{
		Id:        m.Id,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}

	return region
}

func ToDomainListRegion(models []*models.RegionList) []*entity.Region {
	domains := make([]*entity.Region, 0)

	for _, m := range models {
		d := ToDomainRegion(m)
		domains = append(domains, d)
	}

	return domains
}

func ToModelRegion(d *entity.Region) *models.Region {
	region := &models.Region{
		Id:        d.Id,
		Name:      d.Name,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}

	return region
}

func ToModelListRegion(domains []*entity.Region) []*models.Region {
	models := make([]*models.Region, 0)

	for _, d := range domains {
		m := ToModelRegion(d)
		models = append(models, m)
	}

	return models
}

func DataDbqRegion(domain *entity.Region) []interface{} {
	return dbq.Struct(ToModelRegion(domain))
}

func ToDbqStructRegion(domain *entity.Region) (dbqStruct []interface{}) {
	dbqStruct = append(dbqStruct, DataDbqRegion(domain))
	return
}

func ToDbqStructRegionList(domains []*entity.Region) (dbqStruct []interface{}) {
	for _, domain := range domains {
		dbqData := DataDbqRegion(domain)
		dbqStruct = append(dbqStruct, dbqData)
	}
	return
}
