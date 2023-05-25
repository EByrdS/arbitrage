package binance

import (
	"io"
	"net/http"
	"net/url"
)

type IRest interface {
	Delete(endpoint string, payload io.Reader, securityLevel string) (*http.Response, error)
	Get(endpoint string, params url.Values, securityLevel string) (*http.Response, error)
	Post(endpoint string, body io.Reader, securityLevel string) (*http.Response, error)
}

type Binance struct {
	Rest IRest
}
