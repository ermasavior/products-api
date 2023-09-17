package usecase

import "github.com/ProductsAPI/repository"

type Usecase struct {
	Repository repository.RepositoryInterface
}

type NewUsecaseOptions struct {
	Repository repository.RepositoryInterface
}

func NewUsecase(opts NewUsecaseOptions) *Usecase {
	return &Usecase{
		Repository: opts.Repository,
	}
}
