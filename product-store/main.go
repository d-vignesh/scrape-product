package main

import (
	"net/http"
	"os"
	"os/signal"
	"context"
	"time"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/d-vignesh/scrape-product/product-store/utils"
	"github.com/d-vignesh/scrape-product/product-store/handlers"
	"github.com/d-vignesh/scrape-product/product-store/data"
)

const schema = `
		create table if not exists products (
			id  varchar(50) not null,
			url varchar(250) not null,
			name varchar(250) not null,
			imageURL varchar(250) not null,
			description text not null,
			price varchar(20) not null,
			totalReviews int not null
		)
`

func main() {

	logger := utils.NewLogger()

	config := utils.NewConfiguration()

	validator := data.NewValidation()

	db, err := data.NewConnection(config, logger)

	if err != nil {
		logger.Error("unable to connect to db", "error", err)
		panic(err)
	}
	defer db.Close()

	db.MustExec(schema)

	repo := data.NewPostgresRepo(db, logger)

	ph := handlers.NewProductHandler(logger, config, repo, validator)

	sm := mux.NewRouter()

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/store-product", ph.Store)
	postR.Use(ph.MiddlewareValidateProductDetail)

	svr := http.Server {
		Addr:		  config.ServerAddress,
		Handler:	  sm,
		ErrorLog: 	  logger.StandardLogger(&hclog.StandardLoggerOptions{}),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		logger.Info("starting server at address", config.ServerAddress)

		err := svr.ListenAndServe()
		if err != nil {
			logger.Error("could not start the server", "error", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c 
	logger.Info("shutting down the server", "received signal", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	svr.Shutdown(ctx)
}