package main

import (
	"net/http"
	"time"
	"os"
	"os/signal"
	"context"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/d-vignesh/scrape-product/product-scraper/utils"
	"github.com/d-vignesh/scrape-product/product-scraper/services"
	"github.com/d-vignesh/scrape-product/product-scraper/handlers"
)

func main() {

	logger := utils.NewLogger()
	config := utils.NewConfiguration()
	scraper := services.NewScraper(logger)

	sh := handlers.NewScrapingHandler(logger, config, scraper)

	sm := mux.NewRouter()

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/scrape-url", sh.ScrapeURL)

	svr := http.Server {
		Addr :   		config.ServerAddress,
		Handler: 		sm,
		ErrorLog:		logger.StandardLogger(&hclog.StandardLoggerOptions{}),
		ReadTimeout: 	5 * time.Second,
		WriteTimeout:	10 * time.Second,
		IdleTimeout: 	120 * time.Second,
	}

	go func() {
		logger.Info("starting the server at address : ", config.ServerAddress)
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
	logger.Info("shutting down the server", "recieved signal", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	svr.Shutdown(ctx)
}