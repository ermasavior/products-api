package repository

import (
	"context"
	"database/sql"

	"github.com/ProductsAPI/model"
	"github.com/ProductsAPI/utils/sqlwrapper"
)

type RepositoryInterface interface {
	GetDatabase() *sql.DB

	GetProductByProductID(ctx context.Context, productID int) (model.Product, error)
	GetProductVarietiesByProductID(ctx context.Context, productID int) ([]model.ProductVariety, error)

	InsertProduct(ctx context.Context, product model.Product, tx sqlwrapper.Transaction) (int, error)
	InsertProductVarieties(ctx context.Context, productID int, varieties []model.ProductVariety, tx sqlwrapper.Transaction) error
}
