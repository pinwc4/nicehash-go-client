package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/url"
)

//TODO: Move Credentials to another package
type Credentials struct {
	OrgId     string
	ApiKey    string
	ApiSecret string
}

func NewAuthenticator(credentials *Credentials) resty.RequestMiddleware {
	return func(client *resty.Client, request *resty.Request) error {
		separator := string([]byte{0x00})

		requestUrl, err := url.Parse(request.URL)
		if err != nil {
			return err
		}

		message := fmt.Sprintf(
			"%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s",
			credentials.ApiKey,
			separator,
			request.Header.Get("X-Time"),
			separator,
			request.Header.Get("X-Nonce"),
			separator,
			separator,
			credentials.OrgId,
			separator,
			separator,
			request.Method,
			separator,
			requestUrl.EscapedPath(),
			separator,
			request.QueryParam.Encode(),
		)

		mac := hmac.New(sha256.New, []byte(credentials.ApiSecret))
		mac.Write([]byte(message))

		auth := fmt.Sprintf("%s:%s", credentials.ApiKey, hex.EncodeToString(mac.Sum(nil)))

		request.SetHeaders(map[string]string{
			"X-Organization-Id": credentials.OrgId,
			"X-Auth":            auth,
		})

		return nil
	}
}
