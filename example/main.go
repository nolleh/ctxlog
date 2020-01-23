package main

import (
	"ctxlog"
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
	type Request struct {
		Name string `json:"name"`
		Say string `json:"say"`
	}
	var request Request
	c.Bind(&request)
	type Result struct {
		Message string `json:"message"`
	}
	ctxlog.Log().Info(request)
	result := Result{"Hello, World!"}
	json, _ := json.Marshal(result)
	return c.String(http.StatusOK, string(json))
}
