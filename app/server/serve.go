package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"test-api/app/person/delivery"
	"test-api/app/person/repository"
	"test-api/app/person/usecase"

	"test-api/infrastructure/config"
	"test-api/infrastructure/db"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func Serve() {
	// load application configuration
	if err := config.Load(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	// connect to postgres DB
	if err := db.Connect(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	// http server setup
	e := echo.New()


	dbInst := db.Get().DB

	// repository
	personRepo := repository.NewPersonRepository(dbInst)

	// use cases
	personUseCase := usecase.NewPersonUsecase(personRepo)

	// delivery
	delivery.NewPersonHandler(e, personUseCase)

	// start http server
	go func() {
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Get().App.Port)))
	}()

	// graceful shutdown setup
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	logrus.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = e.Shutdown(ctx)
	logrus.Infof("server shutdowns gracefully")
}
