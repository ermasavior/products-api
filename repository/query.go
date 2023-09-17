package repository

const (
	queryGetProductByID = `
		SELECT
			product_id, name, description, rating
		FROM products
		WHERE product_id = $1
	`

	queryGetProductVarietiesByProductID = `
		SELECT
			variety_id, variant, price, stock
		FROM product_varieties
		WHERE product_id = $1
	`
)
