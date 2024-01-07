package office_handler

import (
	"sagara/domain/repository"
	"sagara/domain/usecase"
	"sagara/internal/usecase/office"
	"sagara/pkg/service/jwt"

	"github.com/gorilla/mux"
)

type officeHandler struct {
	jwtService    jwt.JWTService
	officeUseCase usecase.OfficeUseCase
}

func OfficeHandler(
	r *mux.Router,
	jwtService jwt.JWTService,
	officeRepository repository.OfficeRepository,
) {
	officeUseCase := office.NewOfficeInteractor(officeRepository)
	handler := &officeHandler{
		jwtService:    jwtService,
		officeUseCase: officeUseCase,
	}
	r.HandleFunc("/apis/v1/office", handler.GetOffice).Methods("GET")
	r.HandleFunc("/apis/v1/office/{id}", handler.DetailOffice).Methods("GET")
}
