package server

import (
	"context"
	"go-sample/cmd/api/config"
	"go-sample/internal/user"
	"go-sample/internal/user/delivery/http"
	"go-sample/internal/user/entity"
	"go-sample/internal/user/repository"
	"go-sample/internal/user/usecase"
	"reflect"

	v "github.com/core-go/core/v10"
	"github.com/core-go/health"
	log "github.com/core-go/log/zap"
	"github.com/core-go/search/query"
	q "github.com/core-go/sql"
	"github.com/gorilla/mux"
)

type Server struct {
	Health *health.Handler
	User   user.UserHandler
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
	userRepository := repository.NewUserRepository(db)
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
