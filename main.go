package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	_ "sagara/docs"
	"sagara/internal/config"
	docs_handler "sagara/internal/delivery/http/docs"
	office_handler "sagara/internal/delivery/http/office"
	order_handler "sagara/internal/delivery/http/order"
	region_handler "sagara/internal/delivery/http/region"
	static_handler "sagara/internal/delivery/http/static"
	user_handler "sagara/internal/delivery/http/user"
	office_repository "sagara/internal/repository/psql/office"
	order_repository "sagara/internal/repository/psql/order"
	region_repository "sagara/internal/repository/psql/region"
	user_repository "sagara/internal/repository/psql/user"
	"sagara/pkg/logger"
	"sagara/pkg/service/jwt"
	"syscall"

	"github.com/gorilla/mux"
)

var (
	cfg        = config.Server()
	appLogger  = logger.NewApiLogger()
	db         = config.InitDatabase()
	jwtService = jwt.NewJWTService()
	userRepo   = user_repository.NewUserRepository(db)
	regionRepo = region_repository.NewRegionRepository(db)
	officeRepo = office_repository.NewOfficeRepository(db)
	orderRepo  = order_repository.NewOrderRepository(db)
)

func main() {
	psqlConn := config.InitDatabase()
	defer func(db *sql.DB) { _ = db.Close() }(psqlConn)

	router := mux.NewRouter()

	initHandler(router)
	http.Handle("/", router)

	appLogger.Info("sagara Service Run on " + cfg.Port)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		err := http.ListenAndServe(cfg.Port, router)
		if err != nil {
			appLogger.Error(err)
			cancel()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case v := <-quit:
		appLogger.Error(fmt.Sprintf("signal.Notify: %v", v))
	case done := <-ctx.Done():
		appLogger.Error(fmt.Sprintf("ctx.Done: %v", done))
	}
}

func initHandler(router *mux.Router) {
	user_handler.UserHandler(router, jwtService, userRepo)
	region_handler.RegionHandler(router, jwtService, regionRepo)
	office_handler.OfficeHandler(router, jwtService, officeRepo)
	order_handler.OrderHandler(router, jwtService, orderRepo)
	docs_handler.DocsHandler(router)
	static_handler.StaticHandler(router)
}
