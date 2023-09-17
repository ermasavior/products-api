package model

import "github.com/ProductsAPI/generated"

type Product struct {
	ProductID   int              `db:"product_id"`
	Name        string           `validate:"required,max=200,min=3" db:"name"`
	Description string           `validate:"required,max=2000,min=3" db:"description"`
	Rating      float32          `db:"rating"`
	Details     []ProductVariety `validate:"required,dive,required"`
}

type ProductVariety struct {
	VarietyID int     `db:"variety_id"`
	Variant   string  `db:"variant"`
	Price     float32 `validate:"required" db:"price"`
	Stock     int     `validate:"required" db:"stock"`
}

type ValidationProductResult struct {
	IsValid     bool
	Name        string
	Description string
	Details     *ValidationProductDetailResult
}

type ValidationProductDetailResult struct {
	Price string
	Stock string
}

func (v *ValidationProductResult) ToHTTPResponse() *generated.ValidationProductResult {
	res := &generated.ValidationProductResult{}

	if v.Name != "" {
		res.Name = &v.Name
	}
	if v.Description != "" {
		res.Description = &v.Description
	}
	if v.Details != nil {
		res.Details = &generated.ValidationProductDetailResult{
			Price: &v.Details.Price,
			Stock: &v.Details.Stock,
		}
	}

	return res
}
