package repository

import (
	"context"

	"github.com/ProductsAPI/model"
)

type RepositoryInterface interface {
	GetProductByProductID(ctx context.Context, productID int) (model.Product, error)
	GetProductVarietiesByProductID(ctx context.Context, productID int) ([]model.ProductVariety, error)
}
