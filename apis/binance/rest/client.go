package rest

import (
	"net/http"
	"net/url"
)

var client *http.Client = &http.Client{}

func apiURL() url.URL {
	return url.URL{
		Scheme: "https",
		Host:   "api.binance.com",
	}
}
