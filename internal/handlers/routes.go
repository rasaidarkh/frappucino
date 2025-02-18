package handlers

import (
	"database/sql"
	"frappuccino/internal/handlers/middleware"
	"frappuccino/internal/repository"
	"frappuccino/internal/service"
	"log"
	"log/slog"
	"net/http"

	"github.com/go-redis/redis/v8"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", middleware.Middleware(http.NotFoundHandler().ServeHTTP))

	return mux
}

type APIServer struct {
	address string
	mux     *http.ServeMux
	db      *sql.DB
	logger  *slog.Logger
	redis   *redis.Client
}

func NewAPIServer(address string, db *sql.DB, logger *slog.Logger, rdb *redis.Client) *APIServer {
	return &APIServer{
		address: address,
		mux:     Routes(),
		db:      db,
		logger:  logger,
		redis:   rdb,
	}
}

func (s *APIServer) Run() {
	// Logging http server initialization
	s.logger.Info("API server listening on " + s.address)

	// #######################
	// Repository Layer
	// #######################
	inventoryRepository := repository.NewInventoryRepository(s.db)
	userRepository := repository.NewUserRepository(s.db, s.redis)
	//menuRepository := repository.NewMenuRepository(s.db, s.logger)
	//orderRepository := repository.NewOrderRepository(s.db, s.logger)

	// #######################
	// Business Layer
	// #######################
	inventoryService := service.NewInventoryService(inventoryRepository)
	userService := service.NewUserService(userRepository)
	//menuService := service.NewMenuService(menuRepository, s.logger)
	//orderService := service.NewOrderService(orderRepository, s.logger)

	// #######################
	// Presentation Layer
	// #######################
	inventoryHandler := NewInventoryHandler(inventoryService, s.logger)
	userHandler := NewUserHandler(userService, s.logger)
	//menuHandler := handlers.NewMenuHandler(menuService, s.logger)
	//orderHandler := handlers.NewOrderHandler(orderService, s.logger)

	// #######################
	// Registering Endpoints
	// #######################
	inventoryHandler.RegisterEndpoints(s.mux)
	userHandler.RegisterEndpoints(s.mux)
	//menuHandler.RegisterEndpoints(s.mux)
	//orderHandler.RegisterEndpoints(s.mux)

	// #######################
	// Repository Layer
	// #######################
	//repositoryLayer := repository.NewRepository(s.db, s.logger)

	// #######################
	// Business Layer
	// #######################
	//serviceLayer := service.NewService(repositoryLayer, s.logger)

	// #######################
	// Presentation Layer
	// #######################
	//httpLayer := handlers.NewHandler(serviceLayer, s.logger)

	s.logger.Info("API server listening on " + s.address)
	log.Fatal(http.ListenAndServe(s.address, s.mux))
}
