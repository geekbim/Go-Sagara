package office_handler

import (
	"context"
	"net/http"
	"sagara/internal/delivery/response"
	"sagara/pkg/common"
	"sagara/pkg/exceptions"
	"sagara/pkg/utils"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-multierror"
)

// ## DetailOffice
//
//	@Title			Detail Office
//	@Summary		Detail Office
//	@Description	API for get detail office
//	@Tags			Office
//	@Accept			json
//	@Produce		json
//	@Param			Authorization			header	string		true	"Authorization"
//	@Param			id						path	string		false	"Office ID"
//	@Router			/apis/v1/office/{id} [get]
func (handler *officeHandler) DetailOffice(w http.ResponseWriter, r *http.Request) {
	var (
		ctx      = context.Background()
		vars     = mux.Vars(r)
		officeId = vars["id"]
		multierr *multierror.Error
	)

	authHeader := r.Header.Get("Authorization")
	_, err := handler.jwtService.GetUserId(authHeader)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRAUTHORIZED), multierr.Errors)
		return
	}

	officeUuid, _ := common.StringToID(officeId)

	office, errUseCase := handler.officeUseCase.DetailOffice(ctx, officeUuid)
	if errUseCase != nil {
		utils.RespondWithError(w, http.StatusBadRequest, errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response.MapOfficeDomainToResponse(office))
}
