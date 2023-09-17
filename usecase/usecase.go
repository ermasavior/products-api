package usecase

import (
	"github.com/ProductsAPI/repository"
	"github.com/go-playground/validator/v10"
)

type Usecase struct {
	Repository repository.RepositoryInterface
	validate   *validator.Validate
}

type NewUsecaseOptions struct {
	Repository repository.RepositoryInterface
}

func NewUsecase(opts NewUsecaseOptions) *Usecase {
	return &Usecase{
		Repository: opts.Repository,
		validate:   initValidator(),
	}
}
