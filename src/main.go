package main

import (
	application "backend-boilerplate/config"
	log "backend-boilerplate/logger"
	"backend-boilerplate/models"
	"backend-boilerplate/src/api/routes"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title           Swagger Boilerplate API
// @version         1.0
// @host      		localhost:9000
// @BasePath  		/api
func main() {
	if err := application.LoadConfig(); err != nil {
		log.Fatal().Err(err)
	}
	if err := application.InitializeApplication(); err != nil {
		log.Fatal().Err(err).Msg("error initializing application")
	}
	models.SetDatabase(application.DB)
	setupDebug()

	server := &http.Server{
		Addr:    ":9000",
		Handler: routes.SetupRouter(),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("error binding to port")
		}
	}()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt)
	_ = <-osSignals
	log.Info().Msg("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("error shutting down server")
	}
	log.Info().Msg("server shutdown complete")
}

func setupDebug() {
	if application.GetConfig().ApplicationConfig.IsDebug {
		log.Info().Msg("running in mode: debug (app)")
		gin.SetMode(gin.DebugMode)
	} else {
		log.Info().Msg("running in mode: production (app)")
		gin.SetMode(gin.ReleaseMode)
	}
}
