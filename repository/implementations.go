package repository

import (
	"context"
	"database/sql"

	"github.com/ProductsAPI/model"
	"github.com/ProductsAPI/utils/sqlwrapper"
)

func (r *Repository) GetDatabase() *sql.DB {
	return r.Db
}

func (r *Repository) GetProductByProductID(ctx context.Context, productID int) (model.Product, error) {
	var (
		result model.Product
	)

	err := r.Db.QueryRowContext(ctx, queryGetProductByID, productID).
		Scan(&result.ProductID, &result.Name, &result.Description, &result.Rating)

	if err == sql.ErrNoRows {
		return result, nil
	}

	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *Repository) GetProductVarietiesByProductID(ctx context.Context, productID int) ([]model.ProductVariety, error) {
	var (
		result []model.ProductVariety
	)

	rows, err := r.Db.QueryContext(ctx, queryGetProductVarietiesByProductID, productID)
	if rows == nil {
		return result, nil
	}
	defer rows.Close()

	for rows.Next() {
		var row model.ProductVariety

		err := rows.Scan(&row.VarietyID, &row.Variant, &row.Price, &row.Stock)
		if err != nil {
			return []model.ProductVariety{}, err
		}

		result = append(result, row)
	}

	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *Repository) InsertProduct(ctx context.Context, product model.Product, tx sqlwrapper.Transaction) (int, error) {
	var id int

	err := tx.QueryRow(queryInsertNewProduct,
		product.Name, product.Description, product.Rating,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, err
}

func (r *Repository) InsertProductVarieties(ctx context.Context, productID int, varieties []model.ProductVariety, tx sqlwrapper.Transaction) error {
	for _, v := range varieties {
		_, err := tx.ExecContext(ctx, queryInsertNewProductVariety,
			productID, v.Variant, v.Price, v.Stock)

		if err != nil {
			return err
		}
	}

	return nil
}
