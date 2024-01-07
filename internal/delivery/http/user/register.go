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

	"golang.org/x/crypto/bcrypt"
)

// ## Register
//
//	@Title			Register
//	@Summary		Register
//	@Description	API for register
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Payload			body	request.RegisterUser	true	"Request Payload"
//	@Router			/apis/v1/user/register [post]
func (handler *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req request.RegisterUser

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRDOMAIN), []error{err})
		return
	}

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRDOMAIN), []error{err})
		return
	}

	user, errValidate := entity.NewUser(&entity.UserDTO{
		Name:           req.Name,
		Email:          req.Email,
		Password:       string(passwordBytes),
		Position:       req.Position,
		PhoneNumber:    req.PhoneNumber,
		CompanyName:    req.CompanyName,
		CompanyAddress: req.CompanyAddress,
		Information:    req.Information,
	})
	if errValidate != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRDOMAIN), errValidate.Errors)
		return
	}

	ctx := context.Background()
	user, errUseCase := handler.userUseCase.Register(ctx, user)
	if errUseCase != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(exceptions.ERRBUSSINESS), errUseCase.Errors.Errors)
		return
	}

	token := handler.jwtService.GenerateToken(user.Id.String())

	utils.RespondWithJSON(w, http.StatusOK, response.MapUserDomainToResponse(user, token))
}
