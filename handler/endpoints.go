package handler

import (
	"net/http"

	"github.com/ProductsAPI/generated"
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
	return nil
}

func (s *Server) UpdateProduct(c echo.Context, productID int) error {
	return nil
}
