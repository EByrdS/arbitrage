package authentication

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Sign(method string, path string, body string, timestamp int64, secret string) string {
	if method != "GET" && method != "POST" && method != "DELETE" {
		panic(fmt.Sprintf("Unrecognized method %s", method))
	}

	signature_payload := fmt.Sprint(timestamp) + method + path

	if method == "POST" {
		signature_payload += body
	}

	h := hmac.New(sha256.New, []byte(secret))

	_, err := h.Write([]byte(signature_payload))
	if err != nil {
		panic(err)
	}

	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}
