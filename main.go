package main

import (
	"context"

	cfg "github.com/core-go/config"
	"github.com/core-go/core"
	mid "github.com/core-go/log/middleware"
	log "github.com/core-go/log/zap"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"go-service/internal/app"
)

func main() {
	var conf app.Config
	err := cfg.Load(&conf, "./configs/config")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()

	log.Initialize(conf.Log)
	r.Use(mid.BuildContext)
	logger := mid.NewLogger()
	if log.IsInfoEnable() {
		r.Use(mid.Logger(conf.MiddleWare, log.InfoFields, logger))
	}
	r.Use(mid.Recover(log.PanicMsg))

	ctx := context.Background()
	err = app.Route(ctx, r, conf)
	if err != nil {
		panic(err)
	}
	log.Info(ctx, core.ServerInfo(conf.Server))
	server := core.CreateServer(conf.Server, r)
	if err = server.ListenAndServe(); err != nil {
		log.Error(ctx, err.Error())
	}
}
