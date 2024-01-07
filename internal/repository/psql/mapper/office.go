package mapper

import (
	"sagara/domain/entity"
	"sagara/internal/repository/psql/models"
	"sagara/pkg/common"

	"github.com/rocketlaunchr/dbq/v2"
)

func ToDomainOffice(m *models.OfficeList) *entity.Office {
	officeId, _ := common.StringToID(m.Id)
	office := &entity.Office{
		Id:   officeId,
		Name: m.Name,
		Image: &entity.DataObject{
			Path: m.ImagePath,
		},
		RoomType:    m.RoomType,
		Description: m.Description,
		AdditionalData: &entity.OfficeAdditionalData{
			Capacity: m.AdditionalDataCapacity,
			Table:    m.AdditionalDataTable,
			Chair:    m.AdditionalDataChair,
			SpecialOffer: &entity.SpecialOfferAdditionalData{
				Title:       m.AdditionalDataSpecialOfferTitle,
				Description: m.AdditionalDataSpecialOfferDescription,
				Image: &entity.DataObject{
					Path: m.AdditionalDataSpecialOfferImagePath,
				},
			},
			SpecialPrice: &entity.SpecialPriceAdditionalData{
				Image: &entity.DataObject{
					Path: m.AdditionalDataSpecialPriceImagePath,
				},
			},
		},
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}

	return office
}

func ToDomainListOffice(models []*models.OfficeList) []*entity.Office {
	domains := make([]*entity.Office, 0)

	for _, m := range models {
		d := ToDomainOffice(m)
		domains = append(domains, d)
	}

	return domains
}

func ToModelOffice(d *entity.Office) *models.Office {
	office := &models.Office{
		Id:   d.Id.String(),
		Name: d.Name,
		Image: &models.DataObject{
			Id:        d.Image.Id.String(),
			Path:      d.Image.Path,
			CreatedAt: d.Image.CreatedAt,
			UpdatedAt: d.Image.UpdatedAt,
		},
		Region: &models.Region{
			Id:        d.Region.Id,
			Name:      d.Region.Name,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		},
		AdditionalData: &models.OfficeAdditionalData{
			Capacity: d.AdditionalData.Capacity,
			Table:    d.AdditionalData.Table,
			Chair:    d.AdditionalData.Chair,
		},
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}

	if d.AdditionalData.SpecialOffer != nil {
		office.AdditionalData.SpecialOffer.Title = d.AdditionalData.SpecialOffer.Title
		office.AdditionalData.SpecialOffer.Description = d.AdditionalData.SpecialOffer.Description
		if d.AdditionalData.SpecialOffer.Image != nil {
			office.AdditionalData.SpecialOffer.Image.Id = d.AdditionalData.SpecialOffer.Image.Id.String()
			office.AdditionalData.SpecialOffer.Image.Path = d.AdditionalData.SpecialOffer.Image.Path
			office.AdditionalData.SpecialOffer.Image.CreatedAt = d.AdditionalData.SpecialOffer.Image.CreatedAt
			office.AdditionalData.SpecialOffer.Image.UpdatedAt = d.AdditionalData.SpecialOffer.Image.UpdatedAt

		}
	}

	if d.AdditionalData.SpecialPrice != nil {
		if d.AdditionalData.SpecialPrice.Image != nil {
			office.AdditionalData.SpecialPrice.Image.Id = d.AdditionalData.SpecialPrice.Image.Id.String()
			office.AdditionalData.SpecialPrice.Image.Path = d.AdditionalData.SpecialPrice.Image.Path
			office.AdditionalData.SpecialPrice.Image.CreatedAt = d.AdditionalData.SpecialPrice.Image.CreatedAt
			office.AdditionalData.SpecialPrice.Image.UpdatedAt = d.AdditionalData.SpecialPrice.Image.UpdatedAt

		}
	}

	return office
}

func ToModelListOffice(domains []*entity.Office) []*models.Office {
	models := make([]*models.Office, 0)

	for _, d := range domains {
		m := ToModelOffice(d)
		models = append(models, m)
	}

	return models
}

func DataDbqOffice(domain *entity.Office) []interface{} {
	return dbq.Struct(ToModelOffice(domain))
}

func ToDbqStructOffice(domain *entity.Office) (dbqStruct []interface{}) {
	dbqStruct = append(dbqStruct, DataDbqOffice(domain))
	return
}

func ToDbqStructOfficeList(domains []*entity.Office) (dbqStruct []interface{}) {
	for _, domain := range domains {
		dbqData := DataDbqOffice(domain)
		dbqStruct = append(dbqStruct, dbqData)
	}
	return
}
