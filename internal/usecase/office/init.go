package office

import (
	"sagara/domain/repository"
	"sagara/domain/usecase"
)

type officeInteractor struct {
	officeRepository repository.OfficeRepository
}

func NewOfficeInteractor(
	officeRepository repository.OfficeRepository,
) usecase.OfficeUseCase {
	return &officeInteractor{
		officeRepository: officeRepository,
	}
}
