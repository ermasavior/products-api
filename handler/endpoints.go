package handler

import (
	"fmt"
	"net/http"

	"github.com/ProductsAPI/generated"
	"github.com/labstack/echo/v4"
)

func (s *Server) Hello(ctx echo.Context, params generated.HelloParams) error {

	var resp generated.HelloResponse
	resp.Message = fmt.Sprintf("Hello User %d", params.Id)
	return ctx.JSON(http.StatusOK, resp)
}
