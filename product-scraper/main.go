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

	// config holds all the configuration variables needed by the scraping service
	config := utils.NewConfiguration()

	// scraper provides the methods to perform web scraping
	scraper := services.NewScraper(logger)

	// sh(ScrapingHandler) encapsulates all the entities required by Scraper service
	sh := handlers.NewScrapingHandler(logger, config, scraper)

	// create a serve mux
	sm := mux.NewRouter()

	// register handlers on the serve mux for required routes
	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/scrape-url", sh.ScrapeURL)

	// create a new server
	svr := http.Server {
		Addr :   		config.ServerAddress,
		Handler: 		sm,
		ErrorLog:		logger.StandardLogger(&hclog.StandardLoggerOptions{}),
		ReadTimeout: 	5 * time.Second,
		WriteTimeout:	10 * time.Second,
		IdleTimeout: 	120 * time.Second,
	}

	// starting the server in a separate go routine
	go func() {
		logger.Info("starting the server at address : ", config.ServerAddress)
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
	logger.Info("shutting down the server", "recieved signal", sig)
	
	// shutting down the server with a 30 seconds gracetime
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	svr.Shutdown(ctx)
}