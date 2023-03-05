package main

import (
	// "github.com/nolleh/ctxlog"
	"encoding/json"
	"github.com/nolleh/ctxlog/middleware"
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	// Echo instance
	e := echo.New()
	e.Use(middleware.CtxLogger())

	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	// type Request struct {
	// 	Name string `json:"name"`
	// 	Say string `json:"say"`
	// }
	// var request Request
	// c.Bind(&request)
	type Result struct {
		Message string `json:"message"`
	}

	// and if you want to add some additional log, use it as log stream!
	// experience colorized, and pretty json formatting
	// ctxlog.Log().Info(request)

	result := Result{"Hello, World!"}
	json, _ := json.Marshal(result)
	return c.String(http.StatusOK, string(json))
}
