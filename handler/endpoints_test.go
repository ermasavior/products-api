package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ProductsAPI/generated"
	"github.com/ProductsAPI/model"
	"github.com/ProductsAPI/usecase"
	"github.com/bytedance/sonic"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func initHTTPCall(method, path string) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	req := httptest.NewRequest(method, path, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return rec, c
}

func buildHTTPRequestBody(c echo.Context, method, path string, bodyRequest interface{}) echo.Context {
	reqBytes, _ := sonic.Marshal(bodyRequest)
	reqBody := string(reqBytes)
	request, _ := http.NewRequest(method, path, strings.NewReader(reqBody))
	request.Header.Add("Content-Type", "application/json")
	c.SetRequest(request)

	return c
}

func TestGetProduct(t *testing.T) {
	method, path := http.MethodGet, "/products/1"

	ctrl := gomock.NewController(t)
	mockUC := usecase.NewMockUsecaseInterface(ctrl)

	var (
		productID   = 1
		name        = "Dummy"
		description = "Test description"
		rating      = float32(4.5)
		varietyID   = 999
		variant     = ""
		price       = float32(10000)
		stock       = 10
	)

	type args struct {
		productID int
	}
	tests := []struct {
		name           string
		args           args
		mockFunc       func()
		wantStatusCode int
		wantRes        generated.GetProductResponse
	}{
		{
			name: "failed - invalid product ID",
			args: args{
				productID: 0,
			},
			wantStatusCode: http.StatusBadRequest,
			wantRes: generated.GetProductResponse{
				Success: false,
				Error: &generated.ErrorResponse{
					Message: ErrorInvalidProductID,
				},
			},
		},
		{
			name: "failed - usecase returns error",
			args: args{
				productID: productID,
			},
			mockFunc: func() {
				mockUC.EXPECT().GetProduct(gomock.Any(), productID).
					Return(model.Product{}, errors.New("error usecase")).Times(1)
			},
			wantStatusCode: http.StatusInternalServerError,
			wantRes: generated.GetProductResponse{
				Success: false,
				Error: &generated.ErrorResponse{
					Message: "error usecase",
				},
			},
		},
		{
			name: "success",
			args: args{
				productID: productID,
			},
			mockFunc: func() {
				mockUC.EXPECT().GetProduct(gomock.Any(), productID).
					Return(model.Product{
						ProductID:   productID,
						Name:        name,
						Description: description,
						Rating:      rating,
						Details: []model.ProductVariety{
							{
								VarietyID: varietyID,
								Variant:   variant,
								Price:     price,
								Stock:     stock,
							},
						},
					}, nil).Times(1)
			},
			wantStatusCode: http.StatusOK,
			wantRes: generated.GetProductResponse{
				Success: true,
				Data: &generated.Product{
					ProductId:   &productID,
					Name:        name,
					Description: description,
					Rating:      &rating,
					Details: []generated.ProductDetail{
						{
							VarietyId: &varietyID,
							Variant:   &variant,
							Price:     price,
							Stock:     stock,
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{
				Usecase: mockUC,
			}

			if tt.mockFunc != nil {
				tt.mockFunc()
			}

			rec, c := initHTTPCall(method, path)
			s.GetProduct(c, tt.args.productID)

			var got generated.GetProductResponse
			sonic.Unmarshal(rec.Body.Bytes(), &got)

			assert.EqualValues(t, tt.wantStatusCode, rec.Code)
			assert.EqualValues(t, tt.wantRes, got)
		})
	}
}

func TestAddProduct(t *testing.T) {
	method, path := http.MethodPost, "/products"

	ctrl := gomock.NewController(t)
	mockUC := usecase.NewMockUsecaseInterface(ctrl)

	var (
		productID   = 1
		name        = "Dummy"
		description = "Test description"
		rating      = float32(4.5)
		variant     = ""
		price       = float32(10000)
		stock       = 10

		invalidMsg = "invalid msg"
		emptyStr   = ""
	)

	validReq := generated.Product{
		Name:        name,
		Description: description,
		Rating:      &rating,
		Details: []generated.ProductDetail{
			{
				Variant: &variant,
				Price:   price,
				Stock:   stock,
			},
		},
	}

	productInput := model.Product{
		Name:        name,
		Description: description,
		Rating:      rating,
		Details: []model.ProductVariety{
			{
				Variant: variant,
				Price:   price,
				Stock:   stock,
			},
		},
	}

	type args struct {
		req generated.AddProductRequest
	}
	tests := []struct {
		name           string
		args           args
		mockFunc       func()
		wantStatusCode int
		wantRes        generated.AddProductResponse
	}{
		{
			name: "success",
			args: args{
				req: validReq,
			},
			mockFunc: func() {
				mockUC.EXPECT().AddProduct(gomock.Any(), productInput).
					Return(productID, model.ValidationProductResult{
						IsValid: true,
					}, nil).Times(1)
			},
			wantStatusCode: http.StatusOK,
			wantRes: generated.AddProductResponse{
				Success:   true,
				ProductId: &productID,
			},
		},
		{
			name: "failed - error validation",
			args: args{
				req: validReq,
			},
			mockFunc: func() {
				mockUC.EXPECT().AddProduct(gomock.Any(), productInput).
					Return(0, model.ValidationProductResult{
						IsValid: false,
						Name:    invalidMsg,
						Details: &model.ValidationProductDetailResult{
							Stock: invalidMsg,
						},
					}, nil).Times(1)
			},
			wantStatusCode: http.StatusBadRequest,
			wantRes: generated.AddProductResponse{
				Success: false,
				Validation: &generated.ValidationProductResult{
					Name:        &invalidMsg,
					Description: &emptyStr,
					Details: &generated.ValidationProductDetailResult{
						Stock: &invalidMsg,
						Price: &emptyStr,
					},
				},
			},
		},
		{
			name: "failed - error from usecase",
			args: args{
				req: validReq,
			},
			mockFunc: func() {
				mockUC.EXPECT().AddProduct(gomock.Any(), productInput).
					Return(0, model.ValidationProductResult{}, errors.New("error usecase")).Times(1)
			},
			wantStatusCode: http.StatusInternalServerError,
			wantRes: generated.AddProductResponse{
				Success: false,
				Error: &generated.ErrorResponse{
					Message: "error usecase",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{
				Usecase: mockUC,
			}
			if tt.mockFunc != nil {
				tt.mockFunc()
			}

			rec, c := initHTTPCall(method, path)
			c = buildHTTPRequestBody(c, method, path, tt.args.req)
			s.AddProduct(c)

			var got generated.AddProductResponse
			sonic.Unmarshal(rec.Body.Bytes(), &got)

			assert.EqualValues(t, tt.wantStatusCode, rec.Code)
			assert.EqualValues(t, tt.wantRes, got)
		})
	}
}
