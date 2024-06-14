package rest

import (
	_ "github.com/bxcodec/go-clean-arch/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewSwaggerHandler(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
