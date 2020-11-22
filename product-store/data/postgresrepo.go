package data

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/hashicorp/go-hclog"
	uuid "github.com/satori/go.uuid"
)

type PostgresRepo struct {
	db  *sqlx.DB 
	logger hclog.Logger 
}

func NewPostgresRepo(db *sqlx.DB, logger hclog.Logger) *PostgresRepo {
	return &PostgresRepo{db, logger}
}

func (repo *PostgresRepo) Create(ctx context.Context, prodDetail *ProductDetail) error {
	prodDetail.ID = uuid.NewV4().String()
	repo.logger.Debug("creating new product", hclog.Fmt("%#v", prodDetail))
	query := "insert into products (id, url, name, imageURL, description, price, totalReviews) values($1, $2, $3, $4, $5, $6, $7)"
	_, err := repo.db.ExecContext(ctx, query, prodDetail.ID, prodDetail.URL, prodDetail.Product.Name, prodDetail.Product.ImageURL,
								  prodDetail.Product.Description, prodDetail.Product.Price, prodDetail.Product.TotalReviews)
	if err != nil {
		return err
	}
	return nil
}