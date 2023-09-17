package usecase

import (
	"context"

	"github.com/ProductsAPI/model"
)

func (u *Usecase) GetProduct(ctx context.Context, productID int) (model.Product, error) {
	var (
		product model.Product
		err     error
	)

	product, err = u.Repository.GetProductByProductID(ctx, productID)
	if err != nil {
		return model.Product{}, err
	}

	product.Details, err = u.Repository.GetProductVarietiesByProductID(ctx, productID)
	if err != nil {
		return model.Product{}, err
	}

	if err != nil {
		return product, err
	}

	return product, nil
}
