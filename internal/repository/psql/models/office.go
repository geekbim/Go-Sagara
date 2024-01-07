package models

import (
	"time"
)

type Office struct {
	Id             string                `dbq:"id"`
	Name           string                `dbq:"name"`
	RoomType       string                `dbq:"room_type"`
	Description    string                `dbq:"description"`
	Image          *DataObject           `dbq:"image"`
	Region         *Region               `dbq:"region"`
	AdditionalData *OfficeAdditionalData `dbq:"additional_data"`
	CreatedAt      time.Time             `dbq:"created_at"`
	UpdatedAt      time.Time             `dbq:"updated_at"`
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

type OfficeList struct {
	Id                                    string    `dbq:"id"`
	Name                                  string    `dbq:"name"`
	ImagePath                             string    `dbq:"image_path"`
	RoomType                              string    `dbq:"room_type"`
	Description                           string    `dbq:"description"`
	AdditionalDataCapacity                int       `dbq:"additional_data_capacity"`
	AdditionalDataTable                   int       `dbq:"additional_data_table"`
	AdditionalDataChair                   int       `dbq:"additional_data_chair"`
	AdditionalDataSpecialOfferTitle       string    `dbq:"additional_data_special_offer_title"`
	AdditionalDataSpecialOfferDescription string    `dbq:"additional_data_special_offer_description"`
	AdditionalDataSpecialOfferImagePath   string    `dbq:"additional_data_special_offer_image_path"`
	AdditionalDataSpecialPriceImagePath   string    `dbq:"additional_data_special_price_image_path"`
	CreatedAt                             time.Time `dbq:"created_at"`
	UpdatedAt                             time.Time `dbq:"updated_at"`
}

func (Office) TableName() string {
	return "offices"
}

func TableOffice() []string {
	return []string{
		"id",
		"name",
		"room_type",
		"description",
		"image",
		"region",
		"additional_data",
		"created_at",
		"updated_at",
	}
}
