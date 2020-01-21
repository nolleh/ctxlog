package main

import (
	"ctxlog/middleware"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	// Echo instance
	e := echo.New()
	e.Use(middleware.CtxLogger())

	e.POST("/", hello)
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	type Result struct {
		Message string `json:"message"`
	}
	result := Result{"Hello, World!"}
	json, _ := json.Marshal(result)
	return c.String(http.StatusOK, string(json))
}
