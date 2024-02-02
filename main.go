package main

import (
	"fmt"
	"github.com/a-h/templ"
	view "github.com/jafari-mohammad-reza/htmx/views"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	server := echo.New()
	server.GET("/", func(c echo.Context) error {
		return Render(c, view.Index("mohammad reza"))
	})
	server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}

func Render(ctx echo.Context, comp templ.Component) error {
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return comp.Render(ctx.Request().Context(), ctx.Response().Writer)
}
