package data

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/d-vignesh/scrape-product/product-store/utils"
	"github.com/hashicorp/go-hclog"
)

func NewConnection(config *utils.Configuration, logger hclog.Logger) (*sqlx.DB, error) {

	host := config.DBHost
	port := config.DBPort
	dbName := config.DBName 
	user := config.DBUser 
	pass := config.DBPass 
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbName, pass)
	logger.Debug("conn string", conn)

	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		return nil, err
	}
	return db, nil
}