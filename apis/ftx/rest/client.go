package rest

import (
	"net/http"
	"net/url"
)

// for efficiency, there should only be one and reused
var client *http.Client = &http.Client{}

func apiURL() url.URL {
	return url.URL{
		Scheme: "https",
		Host:   "ftx.com",
	}
}
