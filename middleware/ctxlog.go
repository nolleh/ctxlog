package middleware

import (
	"ctxlog"
	"net/http"
	"encoding/json"
	"strings"
	"io/ioutil"
	"bytes"
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

			buf, _ := ioutil.ReadAll(req.Body)
			rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
			rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
			
			var body echo.Map
			ctype := req.Header.Get(echo.HeaderContentType)
			switch {
			case strings.HasPrefix(ctype, echo.MIMEApplicationJSON):
				if err := json.NewDecoder(rdr1).Decode(&body); err != nil {
					// it can be body is null. 
					// TODO: clarify error type
				}
			// TODO: only implemented for json, now. 
			// Change
			}
			
			req.Body = rdr2
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
