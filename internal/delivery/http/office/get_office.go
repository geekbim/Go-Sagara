package office_handler

import (
	"context"
	"net/http"
	"sagara/internal/delivery/request"
	"sagara/internal/delivery/response"
	"sagara/pkg/exceptions"
	"sagara/pkg/utils"
	"strconv"

	"github.com/hashicorp/go-multierror"
)

// ## GetOffice
//
//	@Title			List Office
//	@Summary		List Office
//	@Description	API for get list offices
//	@Tags			Office
//	@Accept			json
//	@Produce		json
//	@Param			Authorization			header	string		true	"Authorization"
//	@Param			page					query	string		false	"Page"
//	@Param			region_id				query	string		false	"Region ID"
//	@Param			limit					query	string		false	"Limit"
//	@Router			/apis/v1/office [get]
func (handler *officeHandler) GetOffice(w http.ResponseWriter, r *http.Request) {
	var (
		ctx      = context.Background()
		query    = r.URL.Query()
		multierr *multierror.Error
	)

	authHeader := r.Header.Get("Authorization")
	_, err := handler.jwtService.GetUserId(authHeader)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRAUTHORIZED), multierr.Errors)
		return
	}

	regionId, _ := strconv.Atoi(query.Get("region_id"))
	limit := utils.StringToInt(query.Get("limit"))
	page := utils.StringToInt(query.Get("page"))

	optionDTO := request.OptionDTO{
		RegionId: regionId,
		Limit:    limit,
		Page:     page,
	}
	payload := request.NewOption(optionDTO)

	offices, count, errUseCase := handler.officeUseCase.ListOffice(ctx, payload)
	if errUseCase != nil {
		utils.RespondWithError(w, http.StatusBadRequest, errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response.MapOfficeListDomainToResponse(offices, count))
}
