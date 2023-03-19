package server

import (
	"context"
	"reflect"

	v "github.com/core-go/core/v10"
	"github.com/core-go/health"
	"github.com/core-go/log/zap"
	"github.com/core-go/search/query"
	q "github.com/core-go/sql"
	"github.com/gorilla/mux"

	"go-service/cmd/api/config"
	"go-service/internal/user"
	"go-service/internal/user/delivery/http"
	"go-service/internal/user/entity"
	"go-service/internal/user/adapter"
	"go-service/internal/user/usecase"
)

type Server struct {
	Health *health.Handler
	User   user.UserPort
}

func NewServer(ctx context.Context, conf config.Config) (*Server, error) {
	db, err := q.OpenByConfig(conf.Sql)
	if err != nil {
		return nil, err
	}
	logError := log.LogError
	validator := v.NewValidator()

	userType := reflect.TypeOf(entity.User{})
	userQueryBuilder := query.NewBuilder(db, "users", userType)
	userSearchBuilder, err := q.NewSearchBuilder(db, userType, userQueryBuilder.BuildQuery)
	if err != nil {
		return nil, err
	}
	userRepository := adapter.NewUserAdapter(db)
	userService := usecase.NewUserService(db, userRepository)
	userHandler := http.NewUserHandler(userSearchBuilder.Search, userService, validator.Validate, logError)

	sqlChecker := q.NewHealthChecker(db)
	healthHandler := health.NewHandler(sqlChecker)

	return &Server{
		Health: healthHandler,
		User:   userHandler,
	}, nil
}

func (s *Server) Run(r *mux.Router) error {
	err := http.UserRoutes(r, s.User)
	if err != nil {
		return err
	}

	return nil
}
