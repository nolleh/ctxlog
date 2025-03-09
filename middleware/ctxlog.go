package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/nolleh/ctxlog"
	"github.com/sirupsen/logrus"
)

func CtxLogger() echo.MiddlewareFunc {
	return CtxLoggerWithLevel(logrus.TraceLevel)
}

func CtxLoggerWithLevel(l logrus.Level) echo.MiddlewareFunc {
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
			rdr3 := io.NopCloser(bytes.NewReader(myWriter.Bytes))
			bind(req, rdr3, &respBody)

			response := Response{&respBody, res.Status}
			data := CtxLogData{request, response, requestId}
			printLogByLevel(l, data)
			return nil
		}
	}
}

func printLogByLevel(l logrus.Level, data CtxLogData) {
	switch l {
	case logrus.TraceLevel:
		ctxlog.Log().Trace(data)
		break
	case logrus.DebugLevel:
		ctxlog.Log().Debug(data)
	case logrus.InfoLevel:
		ctxlog.Log().Info(data)
		break
	case logrus.WarnLevel:
		ctxlog.Log().Warn(data)
		break
	case logrus.ErrorLevel:
		ctxlog.Log().Error(data)
	default:
		ctxlog.Log().Trace(data)
	}
}

func copyBody(req *http.Request) (io.ReadCloser, io.ReadCloser) {
	buf, _ := io.ReadAll(req.Body)
	rdr1 := io.NopCloser(bytes.NewBuffer(buf))
	rdr2 := io.NopCloser(bytes.NewBuffer(buf))
	return rdr1, rdr2
}

func bind(req *http.Request, rdr io.ReadCloser, i any) {
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
