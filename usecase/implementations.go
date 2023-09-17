package usecase

import (
	"context"

	"github.com/ProductsAPI/model"
	"github.com/ProductsAPI/utils/sqlwrapper"
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

func (u *Usecase) AddProduct(ctx context.Context, product model.Product) (int, model.ValidationProductResult, error) {
	var productID int

	err := u.validate.Struct(product)
	if err != nil {
		validateRes := translateError(err)
		return 0, validateRes, nil
	}

	err = sqlwrapper.WithTransaction(ctx, u.Repository.GetDatabase(), func(tx sqlwrapper.Transaction) error {
		var errTx error

		productID, errTx = u.Repository.InsertProduct(ctx, product, tx)
		if errTx != nil {
			return errTx
		}

		errTx = u.Repository.InsertProductVarieties(ctx, productID, product.Details, tx)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, model.ValidationProductResult{}, err
	}

	return productID, model.ValidationProductResult{
		IsValid: true,
	}, nil
}
