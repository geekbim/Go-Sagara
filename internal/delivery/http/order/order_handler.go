package order_handler

import (
	"sagara/domain/repository"
	"sagara/domain/usecase"
	"sagara/internal/usecase/order"
	"sagara/pkg/service/jwt"

	"github.com/gorilla/mux"
)

type orderHandler struct {
	jwtService   jwt.JWTService
	orderUseCase usecase.OrderUseCase
}

func OrderHandler(
	r *mux.Router,
	jwtService jwt.JWTService,
	orderRepository repository.OrderRepository,
) {
	orderUseCase := order.NewOrderInteractor(orderRepository)
	handler := &orderHandler{
		jwtService:   jwtService,
		orderUseCase: orderUseCase,
	}
	r.HandleFunc("/apis/v1/order", handler.CreateOrder).Methods("POST")
}
