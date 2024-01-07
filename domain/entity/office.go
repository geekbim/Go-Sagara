package entity

import (
	"sagara/pkg/common"
	"time"
)

type Office struct {
	Id             common.ID
	Name           string
	RoomType       string
	Description    string
	Image          *DataObject
	Region         *Region
	AdditionalData *OfficeAdditionalData
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type OfficeAdditionalData struct {
	Capacity     int
	Table        int
	Chair        int
	SpecialOffer *SpecialOfferAdditionalData
	SpecialPrice *SpecialPriceAdditionalData
}

type SpecialOfferAdditionalData struct {
	Title       string
	Description string
	Image       *DataObject
}

type SpecialPriceAdditionalData struct {
	Image *DataObject
}
