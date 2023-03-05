module github.com/nolleh/ctxlog

go 1.12

require (
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/nolleh/caption_json_formatter v0.1.0
	github.com/sirupsen/logrus v1.9.0
	golang.org/x/crypto v0.0.0-20220314234659-1baeb1ce4c0b // indirect
	golang.org/x/sys v0.6.0 // indirect
)

replace github.com/nolleh/ctxlog => ../ctxlog
