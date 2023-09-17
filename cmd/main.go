package main

import (
	"os"

	"github.com/ProductsAPI/generated"
	"github.com/ProductsAPI/handler"
	"github.com/ProductsAPI/repository"
	"github.com/ProductsAPI/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	var server generated.ServerInterface = newServer()

	generated.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	dbDsn := os.Getenv("DATABASE_URL")
	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dbDsn,
	})

	var uc usecase.UsecaseInterface = usecase.NewUsecase(usecase.NewUsecaseOptions{
		Repository: repo,
	})

	return handler.NewServer(handler.NewServerOptions{
		Usecase: uc,
	})
}
