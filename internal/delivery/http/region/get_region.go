package region_handler

import (
	"context"
	"net/http"
	"sagara/internal/delivery/request"
	"sagara/internal/delivery/response"
	"sagara/pkg/exceptions"
	"sagara/pkg/utils"

	"github.com/hashicorp/go-multierror"
)

// ## GetRegion
//
//	@Title			List Region
//	@Summary		List Region
//	@Description	API for get list regions
//	@Tags			Region
//	@Accept			json
//	@Produce		json
//	@Param			Authorization			header	string		true	"Authorization"
//	@Param			page					query	string		false	"Page"
//	@Param			key						query	string		false	"Keyword"
//	@Param			limit					query	string		false	"Limit"
//	@Router			/apis/v1/region [get]
func (handler *regionHandler) GetRegion(w http.ResponseWriter, r *http.Request) {
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

	key := query.Get("key")
	limit := utils.StringToInt(query.Get("limit"))
	page := utils.StringToInt(query.Get("page"))

	optionDTO := request.OptionDTO{
		Key:   key,
		Limit: limit,
		Page:  page,
	}
	payload := request.NewOption(optionDTO)

	regions, count, errUseCase := handler.regionUseCase.ListRegion(ctx, payload)
	if errUseCase != nil {
		utils.RespondWithError(w, http.StatusBadRequest, errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response.MapRegionListDomainToResponse(regions, count))
}
