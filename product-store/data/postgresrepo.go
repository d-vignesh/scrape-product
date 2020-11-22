package data

import (
	"context"
	"time"

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

func (repo *PostgresRepo) Create(ctx context.Context, prod *Product) error {
	prod.ID = uuid.NewV4().String()
	prod.CreatedAt = time.Now()
	prod.UpdatedAt = time.Now()
	repo.logger.Debug("creating new product", hclog.Fmt("%#v", prod))
	query := "insert into products (id, url, name, imageURL, description, price, totalReviews, createdAt, UpdatedAt) values($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	_, err := repo.db.ExecContext(ctx, query, prod.ID, prod.URL, prod.Name, prod.ImageURL, prod.Description, 
								  prod.Price, prod.TotalReviews, prod.CreatedAt, prod.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostgresRepo) GetAll(ctx context.Context) ([]*Product, error) {
	prods := make([]*Product, 0)
	rows, err := repo.db.Query("select * from products")
	if err != nil {
		return prods, nil
	}
	for rows.Next() {
		var prod Product
		err = rows.Scan(&prod.ID, &prod.URL, &prod.Name, &prod.ImageURL, &prod.Description,
						&prod.Price, &prod.TotalReviews, &prod.CreatedAt, &prod.UpdatedAt)
		if err != nil {
			repo.logger.Error("error scanning:", err)
			return prods, err
		}
		prods = append(prods, &prod)
	}
	if err = rows.Err(); err != nil {
		repo.logger.Error("error reading rows", err)
		return prods, err
	} 
	return prods, nil
}