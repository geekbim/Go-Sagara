package region_handler

import (
	"sagara/domain/repository"
	"sagara/domain/usecase"
	"sagara/internal/usecase/region"
	"sagara/pkg/service/jwt"

	"github.com/gorilla/mux"
)

type regionHandler struct {
	jwtService    jwt.JWTService
	regionUseCase usecase.RegionUseCase
}

func RegionHandler(
	r *mux.Router,
	jwtService jwt.JWTService,
	regionRepository repository.RegionRepository,
) {
	regionUseCase := region.NewRegionInteractor(regionRepository)
	handler := &regionHandler{
		jwtService:    jwtService,
		regionUseCase: regionUseCase,
	}
	r.HandleFunc("/apis/v1/region", handler.GetRegion).Methods("GET")
}
