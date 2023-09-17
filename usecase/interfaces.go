package usecase

import (
	"context"

	"github.com/ProductsAPI/model"
)

type UsecaseInterface interface {
	GetProduct(ctx context.Context, productID int) (model.Product, error)
}
