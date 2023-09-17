package model

import "errors"

var (
	ErrorInvalidProductID    = errors.New("Invalid Product ID")
	ErrorProductNotFound     = errors.New("Product is not found")
	ErrorEmptyProductDetails = errors.New("Product details cannot be emptied")
)
