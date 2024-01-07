package request

import "sagara/pkg/utils"

type Option struct {
	Key        string
	RegionId   int
	Pagination utils.PaginationOption
}

type OptionDTO struct {
	Key      string
	RegionId int
	Limit    int
	Page     int
}

func NewOption(optionDTO OptionDTO) *Option {
	pagination := utils.NewPagination(optionDTO.Page, optionDTO.Limit)

	return &Option{
		Key:        optionDTO.Key,
		RegionId:   optionDTO.RegionId,
		Pagination: pagination,
	}
}
