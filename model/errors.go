package model

import "errors"

var (
	ErrorProductNotFound     = errors.New("Product is not found")
	ErrorEmptyProductDetails = errors.New("Product details cannot be emptied")
)
