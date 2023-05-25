package authentication

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Sign(body string, timestamp int64, secret string) string {
	signature_payload := body
	if body != "" {
		signature_payload += "&"
	}
	signature_payload += "timestamp=" + fmt.Sprint(timestamp)

	h := hmac.New(sha256.New, []byte(secret))

	_, err := h.Write([]byte(signature_payload))
	if err != nil {
		panic(err)
	}

	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}
