package usecase

import (
	"context"

	"github.com/ProductsAPI/model"
)

type UsecaseInterface interface {
	GetProduct(ctx context.Context, productID int) (model.Product, error)
	AddProduct(ctx context.Context, product model.Product) (int, model.ValidationProductResult, error)
	UpdateProduct(ctx context.Context, input model.Product) (model.ValidationProductResult, error)
}
