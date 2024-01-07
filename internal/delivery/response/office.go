package response

import (
	"fmt"
	"sagara/domain/entity"
	"sagara/internal/config"
)

type OfficeList struct {
	Id             string               `json:"id"`
	Name           string               `json:"name"`
	Image          Image                `json:"image"`
	AdditionalData OfficeAdditionalData `json:"additional_data"`
}

type Office struct {
	Id             string               `json:"id"`
	Name           string               `json:"name"`
	RoomType       string               `json:"room_type"`
	Description    string               `json:"description"`
	Image          Image                `json:"image"`
	AdditionalData OfficeAdditionalData `json:"office_additional_data"`
}

type Image struct {
	Path string `json:"path"`
}

type OfficeAdditionalData struct {
	Capacity     int                         `json:"capacity"`
	Table        int                         `json:"table"`
	Chair        int                         `json:"chair"`
	SpecialOffer *SpecialOfferAdditionalData `json:"special_offer,omitempty"`
	SpecialPrice *SpecialPriceAdditionalData `json:"special_price,omitempty"`
}

type SpecialOfferAdditionalData struct {
	Title       string
	Description string
	Image       *Image
}

type SpecialPriceAdditionalData struct {
	Image *Image
}

type ListOffice struct {
	Offices []*OfficeList `json:"offices"`
	Count   int32         `json:"count"`
}

func MapOfficeDomainToResponseList(office *entity.Office) *OfficeList {
	cfg := config.Server()
	return &OfficeList{
		Id:   office.Id.String(),
		Name: office.Name,
		Image: Image{
			Path: fmt.Sprintf("%s%s", cfg.Port, office.Image.Path),
		},
		AdditionalData: OfficeAdditionalData{
			Capacity: office.AdditionalData.Capacity,
			Table:    office.AdditionalData.Table,
			Chair:    office.AdditionalData.Chair,
		},
	}
}

func MapOfficeListDomainToResponse(offices []*entity.Office, count int32) *ListOffice {
	res := make([]*OfficeList, 0)

	for _, office := range offices {
		res = append(res, MapOfficeDomainToResponseList(office))
	}

	return &ListOffice{
		Offices: res,
		Count:   count,
	}
}

func MapOfficeDomainToResponse(office *entity.Office) *Office {
	cfg := config.Server()
	return &Office{
		Id:          office.Id.String(),
		Name:        office.Name,
		RoomType:    office.RoomType,
		Description: office.Description,
		Image: Image{
			Path: fmt.Sprintf("%s%s", cfg.Port, office.Image.Path),
		},
		AdditionalData: OfficeAdditionalData{
			Capacity: office.AdditionalData.Capacity,
			Table:    office.AdditionalData.Table,
			Chair:    office.AdditionalData.Chair,
			SpecialOffer: &SpecialOfferAdditionalData{
				Title:       office.AdditionalData.SpecialOffer.Title,
				Description: office.AdditionalData.SpecialOffer.Description,
				Image: &Image{
					Path: fmt.Sprintf("%s%s", cfg.Port, office.AdditionalData.SpecialOffer.Image.Path),
				},
			},
			SpecialPrice: &SpecialPriceAdditionalData{
				Image: &Image{
					Path: fmt.Sprintf("%s%s", cfg.Port, office.AdditionalData.SpecialPrice.Image.Path),
				},
			},
		},
	}
}
