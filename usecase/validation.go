package usecase

import (
	"fmt"

	"github.com/ProductsAPI/model"
	"github.com/go-playground/validator/v10"
)

func initValidator() *validator.Validate {
	validate := validator.New()
	return validate
}

func translateError(err error) model.ValidationProductResult {
	if err == nil {
		return model.ValidationProductResult{
			IsValid: true,
		}
	}

	res := model.ValidationProductResult{
		IsValid: false,
	}

	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		errMsg := generateMsgByTagAndParam(e.Tag(), e.Param())

		if e.Field() == "Name" {
			res.Name = errMsg
		}

		if e.Field() == "Description" {
			res.Description = errMsg
		}

		if e.Field() == "Price" {
			if res.Details == nil {
				res.Details = &model.ValidationProductDetailResult{}
			}
			res.Details.Price = errMsg
		}

		if e.Field() == "Stock" {
			if res.Details == nil {
				res.Details = &model.ValidationProductDetailResult{}
			}
			res.Details.Stock = errMsg
		}
	}

	return res
}

func generateMsgByTagAndParam(tag, param string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "max":
		return fmt.Sprintf("Should be less than %v characters", param)
	case "min":
		return fmt.Sprintf("Should be more than %v characters", param)
	}

	return ""
}
