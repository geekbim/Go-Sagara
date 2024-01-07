package user_handler

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

// ## Login
//
//	@Title			Login
//	@Summary		Login
//	@Description	API for login
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Payload			body	request.LoginUser	true	"Request Payload"
//	@Router			/apis/v1/user/login [post]
func (handler *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req request.LoginUser

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRDOMAIN), []error{err})
		return
	}

	user, errValidate := entity.NewUser(&entity.UserDTO{
		Email:    req.Email,
		Password: req.Password,
	})
	if errValidate != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRDOMAIN), errValidate.Errors)
		return
	}

	ctx := context.Background()
	user, err := handler.userUseCase.Login(ctx, user)
	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), err.Errors.Errors)
		return
	}

	token := handler.jwtService.GenerateToken(user.Id.String())

	utils.RespondWithJSON(w, http.StatusOK, response.MapUserDomainToResponse(user, token))
}
