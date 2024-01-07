package region

import (
	"sagara/domain/repository"
	"sagara/domain/usecase"
)

type regionInteractor struct {
	regionRepository repository.RegionRepository
}

func NewRegionInteractor(
	regionRepository repository.RegionRepository,
) usecase.RegionUseCase {
	return &regionInteractor{
		regionRepository: regionRepository,
	}
}
