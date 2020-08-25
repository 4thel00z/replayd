package libreplay

import (
	"net/http"
	"net/url"
)

type Request struct {
	HTTPVersion string      `json:"http_version"`
	Host        string      `json:"host"`
	Method      string      `json:"method"`
	Headers     http.Header `json:"headers"`
	Body        []byte      `json:"body"`
	Form        url.Values  `json:"form"`
}
type StringWriterCloser interface {
	WriteString(s string) (n int, err error)
	Close() error
}
