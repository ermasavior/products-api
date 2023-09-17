package handler

import (
	"io"
	"net/http"

	"github.com/ProductsAPI/generated"
	"github.com/bytedance/sonic"
	"github.com/labstack/echo/v4"
)

func (s *Server) GetProduct(c echo.Context, productID int) error {
	var (
		result generated.GetProductResponse
	)

	if productID <= 0 {
		result.Error = &generated.ErrorResponse{
			Message: ErrorInvalidProductID,
		}
		return c.JSON(http.StatusBadRequest, result)
	}

	ctx := c.Request().Context()

	product, err := s.Usecase.GetProduct(ctx, productID)
	if err != nil {
		result.Error = &generated.ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, result)
	}

	detailsRes := make([]generated.ProductDetail, 0, len(product.Details))
	for _, p := range product.Details {
		varietyID, variant := p.VarietyID, p.Variant
		detailsRes = append(detailsRes, generated.ProductDetail{
			VarietyId: &varietyID,
			Variant:   &variant,
			Price:     p.Price,
			Stock:     p.Stock,
		})
	}
	result = generated.GetProductResponse{
		Success: true,
		Data: &generated.Product{
			ProductId:   &product.ProductID,
			Name:        product.Name,
			Description: product.Description,
			Rating:      &product.Rating,
			Details:     detailsRes,
		},
	}

	return c.JSON(http.StatusOK, result)
}

func (s *Server) AddProduct(c echo.Context) error {
	var (
		result     generated.AddProductResponse
		productReq generated.AddProductJSONRequestBody
	)

	reqBody, _ := io.ReadAll(c.Request().Body)
	_ = sonic.Unmarshal(reqBody, &productReq)

	ctx := c.Request().Context()
	productInput := buildProductInput(productReq)

	productID, validateRes, err := s.Usecase.AddProduct(ctx, productInput)
	if err != nil {
		result.Error = &generated.ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, result)
	}

	if !validateRes.IsValid {
		var validateDetails *generated.ValidationProductDetailResult
		if validateRes.Details != nil {
			validateDetails = &generated.ValidationProductDetailResult{
				Price: &validateRes.Details.Price,
				Stock: &validateRes.Details.Stock,
			}
		}
		result.Validation = &generated.ValidationProductResult{
			Name:        &validateRes.Name,
			Description: &validateRes.Description,
			Details:     validateDetails,
		}
		return c.JSON(http.StatusBadRequest, result)
	}

	result = generated.AddProductResponse{
		Success:   true,
		ProductId: &productID,
	}
	return c.JSON(http.StatusOK, result)
}

func (s *Server) UpdateProduct(c echo.Context, productID int) error {
	var (
		result     generated.UpdateProductResponse
		productReq generated.UpdateProductJSONRequestBody
	)

	if productID <= 0 {
		result.Error = &generated.ErrorResponse{
			Message: ErrorInvalidProductID,
		}
		return c.JSON(http.StatusBadRequest, result)
	}

	reqBody, _ := io.ReadAll(c.Request().Body)
	_ = sonic.Unmarshal(reqBody, &productReq)

	ctx := c.Request().Context()

	productReq.ProductId = &productID
	productInput := buildProductInput(productReq)

	validateRes, err := s.Usecase.UpdateProduct(ctx, productInput)
	if err != nil {
		result.Error = &generated.ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, result)
	}

	if !validateRes.IsValid {
		var validateDetails *generated.ValidationProductDetailResult
		if validateRes.Details != nil {
			validateDetails = &generated.ValidationProductDetailResult{
				Price: &validateRes.Details.Price,
				Stock: &validateRes.Details.Stock,
			}
		}
		result.Validation = &generated.ValidationProductResult{
			Name:        &validateRes.Name,
			Description: &validateRes.Description,
			Details:     validateDetails,
		}
		return c.JSON(http.StatusBadRequest, result)
	}

	result = generated.UpdateProductResponse{
		Success: true,
	}
	return c.JSON(http.StatusOK, result)
}
