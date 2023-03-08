package main

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"

	cjf "github.com/nolleh/caption_json_formatter"
	"github.com/nolleh/ctxlog"
	"github.com/nolleh/ctxlog/middleware"
)

func main() {
	// [optional] retreive ctxlogs formatter (CaptionJsonFormatter)
	formatter := ctxlog.CaptionJsonFormatter()
	// [optional] you can modify format configruation from default.
	formatter.PrettyPrint = true
	// [optional] or, you reset your own formatter
	formatter = cjf.Json()
	// [optional] and set any formatter compatiable with logrus.formatter
	ctxlog.Logger().SetFormatter(formatter)

	// Echo instance
	e := echo.New()
	// [required] use ctxlog
	e.Use(middleware.CtxLogger())
	e.Use(middleware.CtxLoggerWithLevel(logrus.WarnLevel))

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
