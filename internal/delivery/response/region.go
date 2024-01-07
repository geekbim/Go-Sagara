package response

import (
	"sagara/domain/entity"
)

type RegionList struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ListRegion struct {
	Regions []*RegionList `json:"regions"`
	Count   int32         `json:"count"`
}

func MapRegionDomainToResponseList(region *entity.Region) *RegionList {
	return &RegionList{
		Id:   region.Id,
		Name: region.Name,
	}
}

func MapRegionListDomainToResponse(regions []*entity.Region, count int32) *ListRegion {
	res := make([]*RegionList, 0)

	for _, region := range regions {
		res = append(res, MapRegionDomainToResponseList(region))
	}

	return &ListRegion{
		Regions: res,
		Count:   count,
	}
}
