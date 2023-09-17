package handler

import (
	"io"
	"net/http"

	"github.com/ProductsAPI/generated"
	"github.com/ProductsAPI/model"
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
		resp       generated.AddProductResponse
		productReq generated.AddProductJSONRequestBody
	)

	reqBody, _ := io.ReadAll(c.Request().Body)
	_ = sonic.Unmarshal(reqBody, &productReq)

	productDetailsInput := make([]model.ProductVariety, 0, len(productReq.Details))
	for _, p := range productReq.Details {
		var variant string
		if p.Variant != nil {
			variant = *p.Variant
		}

		productDetailsInput = append(productDetailsInput, model.ProductVariety{
			Variant: variant,
			Price:   p.Price,
			Stock:   p.Stock,
		})
	}

	var rating float32
	if productReq.Rating != nil {
		rating = *productReq.Rating
	}

	productInput := model.Product{
		Name:        productReq.Name,
		Description: productReq.Description,
		Rating:      rating,
		Details:     productDetailsInput,
	}

	ctx := c.Request().Context()
	productID, validateRes, err := s.Usecase.AddProduct(ctx, productInput)
	if err != nil {
		resp.Error = &generated.ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	if !validateRes.IsValid {
		var validateDetails *generated.ValidationProductDetailResult
		if validateRes.Details != nil {
			validateDetails = &generated.ValidationProductDetailResult{
				Price: &validateRes.Details.Price,
				Stock: &validateRes.Details.Stock,
			}
		}
		resp.Validation = &generated.ValidationProductResult{
			Name:        &validateRes.Name,
			Description: &validateRes.Description,
			Details:     validateDetails,
		}
		return c.JSON(http.StatusBadRequest, resp)
	}

	resp = generated.AddProductResponse{
		Success:   true,
		ProductId: &productID,
	}
	return c.JSON(http.StatusOK, resp)
}

func (s *Server) UpdateProduct(c echo.Context, productID int) error {
	return nil
}
