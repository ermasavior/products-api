package repository

import (
	"context"
	"database/sql"

	"github.com/ProductsAPI/model"
)

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
