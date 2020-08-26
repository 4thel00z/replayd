package libreplay

import (
	"encoding/json"
	"github.com/monzo/typhon"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Request struct {
	HTTPVersion string      `json:"http_version"`
	URL         string      `json:"url"`
	Method      string      `json:"method"`
	Headers     http.Header `json:"headers"`
	Body        []byte      `json:"body"`
	Form        url.Values  `json:"form"`
}
type StringWriterCloser interface {
	WriteString(s string) (n int, err error)
	Close() error
}

func ParseConfig(path string) (config Config, err error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return Config{}, err
	}
	err = json.Unmarshal(content, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

type Config struct {
	//TODO: add more fields here if you want to make the app more configurable
	Verbose bool   `json:"verbose"`
	Path    string `json:"path"`
}

type Service func(app App) typhon.Service

type Route struct {
	Path        string `json:"path"`
	Method      string `json:"method"`
	CurlExample string `json:"curl_example"`
	longPath    string
}
type Module interface {
	Version() string
	Namespace() string
	Routes() []Route
	HandlerById(int) Service
}
