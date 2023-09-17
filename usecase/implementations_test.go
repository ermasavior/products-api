package usecase

import (
	"context"
	"testing"

	"github.com/ProductsAPI/model"
	"github.com/ProductsAPI/repository"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockRepositoryInterface(ctrl)

	var (
		ctx = context.Background()

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
		name     string
		args     args
		mockFunc func()
		wantErr  bool
		wantRes  model.Product
	}{
		{
			name: "success",
			args: args{
				productID: productID,
			},
			mockFunc: func() {
				mockRepo.EXPECT().GetProductByProductID(gomock.Any(), productID).
					Return(model.Product{
						ProductID:   productID,
						Name:        name,
						Description: description,
						Rating:      rating,
					}, nil).Times(1)
				mockRepo.EXPECT().GetProductVarietiesByProductID(gomock.Any(), productID).
					Return([]model.ProductVariety{
						{
							VarietyID: varietyID,
							Variant:   variant,
							Price:     price,
							Stock:     stock,
						},
					}, nil)
			},
			wantErr: false,
			wantRes: model.Product{
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
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := Usecase{
				Repository: mockRepo,
			}

			if tt.mockFunc != nil {
				tt.mockFunc()
			}

			got, err := uc.GetProduct(ctx, tt.args.productID)

			assert.EqualValues(t, tt.wantRes, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
