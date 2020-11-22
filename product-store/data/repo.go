package data

import (
	"context"
)


// Interface to be satisfied by all repository implementations of Store service
type Repository interface {
	Create(ctx context.Context, product *Product) error
	GetAll(ctx context.Context) ([]*Product, error)
}