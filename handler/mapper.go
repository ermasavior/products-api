package handler

import (
	"github.com/ProductsAPI/generated"
	"github.com/ProductsAPI/model"
)

func buildProductInput(input generated.Product) model.Product {
	var productID int
	if input.ProductId != nil {
		productID = *input.ProductId
	}

	productDetailsInput := make([]model.ProductVariety, 0, len(input.Details))
	for _, p := range input.Details {
		var varietyID int
		if p.VarietyId != nil {
			varietyID = *p.VarietyId
		}

		var variant string
		if p.Variant != nil {
			variant = *p.Variant
		}

		productDetailsInput = append(productDetailsInput, model.ProductVariety{
			VarietyID: varietyID,
			Variant:   variant,
			Price:     p.Price,
			Stock:     p.Stock,
		})
	}

	var rating float32
	if input.Rating != nil {
		rating = *input.Rating
	}

	return model.Product{
		ProductID:   productID,
		Name:        input.Name,
		Description: input.Description,
		Rating:      rating,
		Details:     productDetailsInput,
	}
}
