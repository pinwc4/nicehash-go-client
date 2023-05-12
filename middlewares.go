package nhclient

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

var defaultHeaders = func(c *resty.Client, request *resty.Request) error {
	request.SetHeaders(map[string]string{
		"X-Time":       strconv.FormatInt(time.Now().UnixMilli(), 10),
		"X-Nonce":      uuid.NewString(),
		"X-Request-Id": uuid.NewString(),
		"Content-Type": "application/json",
	})

	return nil
}

func authenticator(orgId, apiKey, secretKey string) resty.RequestMiddleware {
	return func(client *resty.Client, request *resty.Request) error {
		separator := string([]byte{0x00})

		requestUrl, err := url.Parse(request.URL)
		if err != nil {
			return err
		}

		message := fmt.Sprintf(
			"%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s",
			apiKey,
			separator,
			request.Header.Get("X-Time"),
			separator,
			request.Header.Get("X-Nonce"),
			separator,
			separator,
			orgId,
			separator,
			separator,
			request.Method,
			separator,
			requestUrl.EscapedPath(),
			separator,
			request.QueryParam.Encode(),
		)

		if request.Body != nil {
			message = fmt.Sprintf("%s%s%s", message, separator, request.Body)
		}

		mac := hmac.New(sha256.New, []byte(secretKey))
		mac.Write([]byte(message))

		auth := fmt.Sprintf("%s:%s", apiKey, hex.EncodeToString(mac.Sum(nil)))

		request.SetHeaders(map[string]string{
			"X-Organization-Id": orgId,
			"X-Auth":            auth,
		})

		return nil
	}
}
