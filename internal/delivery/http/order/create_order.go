package order_handler

import (
	"context"
	"encoding/json"
	"net/http"
	"sagara/domain/entity"
	"sagara/internal/delivery/request"
	"sagara/internal/delivery/response"
	"sagara/pkg/exceptions"
	"sagara/pkg/utils"
)

// ## CreateOrder
//
//	@Title			Create Order
//	@Summary		Create Order
//	@Description	API for create order
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			Authorization			header	string		true	"Authorization"
//	@Param			Payload			body	request.CreateOrder	true	"Request Payload"
//	@Router			/apis/v1/order [post]
func (handler *orderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req request.CreateOrder

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRDOMAIN), []error{err})
		return
	}

	order, errValidate := entity.NewOrder(&entity.OrderDTO{
		PicName:          req.PicName,
		DateStart:        req.DateStart,
		DateEnd:          req.DateEnd,
		Payment:          req.Payment,
		CapacityEmployee: req.CapacityEmployee,
		Requirement:      req.Requirement,
		Other:            req.Other,
	})
	if errValidate != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRDOMAIN), errValidate.Errors)
		return
	}

	ctx := context.Background()
	order, errUseCase := handler.orderUseCase.CreateOrder(ctx, order)
	if errUseCase != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errUseCase.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response.MapOrderDomainToResponse(order))
}
