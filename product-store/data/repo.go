package data

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, product *ProductDetail) error
}