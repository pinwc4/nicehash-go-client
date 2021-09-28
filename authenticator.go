package nhclient

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/url"
)

func authenticatorMiddleware(credentials *Credentials) resty.RequestMiddleware {
	return func(c *resty.Client, request *resty.Request) error {
		separator := string([]byte{0x00})

		requestUrl, err := url.Parse(request.URL)
		if err != nil {
			return err
		}

		message := fmt.Sprintf(
			"%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s",
			credentials.apiKey,
			separator,
			request.Header.Get("X-Time"),
			separator,
			request.Header.Get("X-Nonce"),
			separator,
			separator,
			credentials.orgId,
			separator,
			separator,
			request.Method,
			separator,
			requestUrl.EscapedPath(),
			separator,
			request.QueryParam.Encode(),
		)

		mac := hmac.New(sha256.New, []byte(credentials.apiSecret))
		mac.Write([]byte(message))

		auth := fmt.Sprintf("%s:%s", credentials.apiKey, hex.EncodeToString(mac.Sum(nil)))

		request.SetHeaders(map[string]string{
			"X-Organization-Id": credentials.orgId,
			"X-Auth":            auth,
		})

		return nil
	}
}
