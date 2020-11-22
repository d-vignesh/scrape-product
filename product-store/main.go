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

// schema for the products table
const schema = `
		create table if not exists products (
			id  varchar(50) not null,
			url varchar(250) not null,
			name varchar(250) not null,
			imageURL varchar(250) not null,
			description text not null,
			price varchar(20) not null,
			totalReviews int not null,
			createdAt timestamp not null,
			updatedAt timestamp not null,
			primary key (id)
		)
`

func main() {

	logger := utils.NewLogger()

	// config holds all the configuration variables needed by the store service
	config := utils.NewConfiguration()

	// validator contains all the methods that are need to validate the product json in request
	validator := data.NewValidation()

	db, err := data.NewConnection(config, logger)

	if err != nil {
		logger.Error("unable to connect to db", "error", err)
		panic(err)
	}
	defer db.Close()

	db.MustExec(schema)

	// repo contains all the methods that interact with DB to perform CURD operations for product.
	repo := data.NewPostgresRepo(db, logger)

	// (ph)ProductHandler encapsulates all the entities required by Store service
	ph := handlers.NewProductHandler(logger, config, repo, validator)

	// create a serve mux
	sm := mux.NewRouter()

	// register handlers on the serve mux for required routes
	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/store-product", ph.StoreProduct)
	postR.Use(ph.MiddlewareValidateProductDetail)

	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/get-products", ph.GetProducts)

	// create a new server
	svr := http.Server {
		Addr:		  config.ServerAddress,
		Handler:	  sm,
		ErrorLog: 	  logger.StandardLogger(&hclog.StandardLoggerOptions{}),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// starting the server in a separate go routine
	go func() {
		logger.Info("starting server at address", config.ServerAddress)

		err := svr.ListenAndServe()
		if err != nil {
			logger.Error("could not start the server", "error", err)
			os.Exit(1)
		}
	}()

	// create a channel to listen for signals to kill server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// blocking until any signal is received in the channel
	sig := <-c 
	logger.Info("shutting down the server", "received signal", sig)

	// shutting down the server with a 30 seconds gracetime
	ctx, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	svr.Shutdown(ctx)
}