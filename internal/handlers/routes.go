package handlers

import (
	"context"
	"database/sql"
	"frappuccino/internal/repository"
	"frappuccino/internal/service"
	"log"
	"log/slog"
	"net/http"
)

type APIServer struct {
	address string
	mux     *http.ServeMux
	db      *sql.DB
	logger  *slog.Logger
	ctx     context.Context
}

func NewAPIServer(address string, mux *http.ServeMux, db *sql.DB, logger *slog.Logger, ctx context.Context) *APIServer {
	return &APIServer{
		address: address,
		mux:     mux,
		db:      db,
		logger:  logger,
		ctx:     ctx,
	}
}

func (s *APIServer) Run() {
	s.logger.Info("API server listening on " + s.address)

	//Creating Repository Layer

	inventoryRepository := repository.NewInventoryRepository(s.db, s.logger)
	//menuRepository := repository.NewMenuRepository(s.db, s.logger)
	//orderRepository := repository.NewOrderRepository(s.db, s.logger)

	//Creating Business Layer
	inventoryService := service.NewInventoryService(inventoryRepository, s.logger)
	//menuService := service.NewMenuService(menuRepository, s.logger)
	//orderService := service.NewOrderService(orderRepository, s.logger)
	//Creating Presentation Layer
	inventoryHandler := NewInventoryHandler(inventoryService, s.logger)
	//menuHandler := handlers.NewMenuHandler(menuService, s.logger)
	//orderHandler := handlers.NewOrderHandler(orderService, s.logger)

	//Registering Endpoints
	inventoryHandler.RegisterEndpoints(s.mux)
	//menuHandler.RegisterEndpoints(s.mux)
	//orderHandler.RegisterEndpoints(s.mux)

	//Creating Repository Layer
	//repositoryLayer := repository.NewRepository(s.db, s.logger)
	//Creating Business Layer
	//serviceLayer := service.NewService(repositoryLayer, s.logger)
	//Creating Presentation Layer
	//httpLayer := handlers.NewHandler(serviceLayer, s.logger)

	log.Fatal(http.ListenAndServe(s.address, s.mux))
}
