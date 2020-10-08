package middleware

import (
	"bytes"
	"github.com/nolleh/ctxlog"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

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

			rdr1, rdr2 := copyBody(req)
			var body echo.Map
			bind(req, rdr1, &body)
			req.Body = rdr2

			request := Request{req.Header, req.RequestURI, req.Method, &body}

			myWriter := &Writer{res.Writer, nil}
			res.Writer = myWriter
			if err := next(c); err != nil {
				return err
			}

			var respBody echo.Map
			rdr3 := ioutil.NopCloser(bytes.NewReader(myWriter.Bytes))
			bind(req, rdr3, &respBody)

			response := Response{&respBody, res.Status}
			data := CtxLogData{request, response, requestId}
			ctxlog.Log().Trace(data)
			return nil
		}
	}
}

func copyBody(req *http.Request) (io.ReadCloser, io.ReadCloser) {
	buf, _ := ioutil.ReadAll(req.Body)
	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
	return rdr1, rdr2
}

func bind(req *http.Request, rdr io.ReadCloser, i interface{}) {
	ctype := req.Header.Get(echo.HeaderContentType)
	switch {
	case strings.HasPrefix(ctype, echo.MIMEApplicationJSON):
		if err := json.NewDecoder(rdr).Decode(&i); err != nil {
			// it can be body is null.
			// TODO: clarify error type
		}
		// TODO: only implemented for json, now.
		// Change
	}
}

type Writer struct {
	http.ResponseWriter
	Bytes []byte
}

func (w *Writer) Write(bytes []byte) (int, error) {
	w.Bytes = make([]byte, len(bytes))
	copy(w.Bytes, bytes)
	return w.ResponseWriter.Write(bytes)
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
	RequestId string   `json:"requestId"`
}
