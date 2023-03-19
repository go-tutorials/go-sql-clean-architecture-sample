package main

import (
	"context"
	"fmt"

	cfg "github.com/core-go/config"
	"github.com/core-go/core"
	mid "github.com/core-go/log/middleware"
	log "github.com/core-go/log/zap"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"go-sample/cmd/api/config"
	"go-sample/cmd/api/server"
)

func main() {
	var conf config.Config
	err := cfg.Load(&conf, "../../configs/config")
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

	s, err := server.NewServer(context.Background(), conf)
	if err != nil {
		panic(err)
	}

	if err = s.Run(r); err != nil {
		panic(err)
	}
	fmt.Println(core.ServerInfo(conf.Server))

	server := core.CreateServer(conf.Server, r)
	if err = server.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}
}
