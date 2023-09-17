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
	var (
		productID   int
		validateRes model.ValidationProductResult
	)

	err := u.validate.Struct(product)
	if err != nil {
		validateRes = translateError(err)
		return 0, validateRes, nil
	}

	validateRes.IsValid = true

	err = sqlwrapper.WithTransaction(ctx, u.Repository.GetDatabase(), func(tx sqlwrapper.Transaction) error {
		var errTx error

		productID, errTx = u.Repository.InsertProduct(ctx, product, tx)
		if errTx != nil {
			return errTx
		}

		for _, v := range product.Details {
			v.ProductID = productID
			errTx = u.Repository.InsertProductVariety(ctx, v, tx)
			if errTx != nil {
				return errTx
			}
		}

		return nil
	})

	if err != nil {
		return 0, validateRes, err
	}

	return productID, validateRes, nil
}

func (u *Usecase) UpdateProduct(ctx context.Context, input model.Product) (model.ValidationProductResult, error) {
	var validateRes model.ValidationProductResult

	err := u.validate.Struct(input)
	if err != nil {
		validateRes = translateError(err)
		return validateRes, nil
	}

	validateRes.IsValid = true

	currProduct, err := u.GetProduct(ctx, input.ProductID)
	if err != nil {
		return validateRes, err
	}

	if currProduct.ProductID == 0 {
		return validateRes, model.ErrorProductNotFound
	}

	var (
		mapCurrProductVarieties = make(map[int]model.ProductVariety)
		mapInputProductDetails  = make(map[int]bool)

		insertProductVarieties = make([]model.ProductVariety, 0)
		updateProductVarieties = make([]model.ProductVariety, 0)
		deleteProductVarieties = make([]model.ProductVariety, 0)
	)

	for _, v := range currProduct.Details {
		mapCurrProductVarieties[v.VarietyID] = v
	}

	for _, in := range input.Details {
		mapInputProductDetails[in.VarietyID] = true
	}

	for _, in := range input.Details {
		if in.VarietyID == 0 {
			insertProductVarieties = append(insertProductVarieties, in)
			continue
		}

		if _, ok := mapCurrProductVarieties[in.VarietyID]; ok {
			updateProductVarieties = append(updateProductVarieties, in)
			continue
		}
	}

	deleteCountRow := 0
	for _, v := range mapCurrProductVarieties {
		if mapInputProductDetails[v.VarietyID] {
			continue
		}

		deleteProductVarieties = append(deleteProductVarieties, v)
		deleteCountRow += 1
	}

	if len(insertProductVarieties) == 0 && len(updateProductVarieties) == 0 && deleteCountRow == len(mapCurrProductVarieties) {
		return validateRes, model.ErrorEmptyProductDetails
	}

	err = sqlwrapper.WithTransaction(ctx, u.Repository.GetDatabase(), func(tx sqlwrapper.Transaction) error {
		var errTx error

		errTx = u.Repository.UpdateProduct(ctx, input, tx)
		if errTx != nil {
			return errTx
		}

		for _, v := range insertProductVarieties {
			v.ProductID = currProduct.ProductID
			errTx = u.Repository.InsertProductVariety(ctx, v, tx)
			if errTx != nil {
				return errTx
			}
		}

		for _, v := range updateProductVarieties {
			v.ProductID = currProduct.ProductID
			errTx = u.Repository.UpdateProductVariety(ctx, v, tx)
			if errTx != nil {
				return errTx
			}
		}

		for _, v := range deleteProductVarieties {
			v.ProductID = currProduct.ProductID
			errTx = u.Repository.DeleteProductVariety(ctx, v, tx)
			if errTx != nil {
				return errTx
			}
		}

		return nil
	})

	if err != nil {
		return validateRes, err
	}

	return validateRes, nil
}

func (u *Usecase) DeleteProduct(ctx context.Context, productID int) error {
	varieties, err := u.Repository.GetProductVarietiesByProductID(ctx, productID)
	if err != nil {
		return err
	}

	err = sqlwrapper.WithTransaction(ctx, u.Repository.GetDatabase(), func(tx sqlwrapper.Transaction) error {
		var errTx error

		for _, v := range varieties {
			v.ProductID = productID
			errTx = u.Repository.DeleteProductVariety(ctx, v, tx)
			if errTx != nil {
				return errTx
			}
		}

		err = u.Repository.DeleteProduct(ctx, productID, tx)
		if err != nil {
			return err
		}

		return nil
	})

	return nil
}
