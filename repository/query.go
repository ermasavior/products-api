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

	queryUpdateProduct = `
		UPDATE
			products
		SET
			name = $1,
			description = $2,
			rating = $3
		WHERE
			product_id = $4
	`

	queryUpdateProductVariety = `
		UPDATE
			product_varieties
		SET
			variant = $1,
			price = $2,
			stock = $3
		WHERE
			variety_id = $4 AND product_id = $5
	`

	queryDeleteProductVariety = `
		DELETE FROM
			product_varieties
		WHERE
			variety_id = $1 AND product_id = $2
	`

	queryDeleteProduct = `
		DELETE FROM products
		WHERE product_id = $1
	`
)
