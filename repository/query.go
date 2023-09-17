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

	queryInsertNewProduct = `
		INSERT INTO products(name, description, rating)
		VALUES($1, $2, $3)
		RETURNING product_id
	`

	queryInsertNewProductVariety = `
		INSERT INTO product_varieties(product_id, variant, price, stock)
		VALUES($1, $2, $3, $4)
	`
)
