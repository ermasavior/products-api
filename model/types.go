package model

type Product struct {
	ProductID   int     `db:"product_id"`
	Name        string  `db:"name"`
	Description string  `db:"description"`
	Rating      float32 `db:"rating"`
	Details     []ProductVariety
}

type ProductVariety struct {
	VarietyID int     `db:"variety_id"`
	Variant   string  `db:"variant"`
	Price     float32 `db:"price"`
	Stock     int     `db:"stock"`
}
