package middleware

import (
	"ctxlog"
	"net/http"

	"github.com/labstack/echo"
)

func CtxLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			requestId := req.Header.Get(echo.HeaderXRequestID)
			if requestId == "" {
				requestId = res.Header().Get(echo.HeaderXRequestID)
			}

			var body echo.Map
			c.Bind(&body)
			request := Request{req.Header, req.RequestURI, req.Method, &body}

			if err := next(c); err != nil {
				return err
			}

			response := Response{} // TODO
			data := CtxLogData{request, response, requestId}
			ctxlog.Log().Trace(data)
			return nil
		}
	}

}

type Request struct {
	Header http.Header `json:"header"`
	Uri    string      `json:"uri"`
	Method string      `json:"method"`
	Body   *echo.Map   `json:"body"`
}

type Response struct {
	Body   *echo.Map `json:"body"`
	Status int       `json:"status"`
}

type CtxLogData struct {
	Request   Request  `json:"request"`
	Response  Response `json:"response"`
	RequestId string   `json:"requestId`
}
