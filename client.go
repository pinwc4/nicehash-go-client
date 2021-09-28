package nhclient

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"net/url"
	"strconv"
	"time"
)

type Credentials struct {
	orgId     string
	apiKey    string
	apiSecret string
}

type client struct {
	credentials *Credentials

	Mining *mining
}

func New(credentials *Credentials) *client {
	var defaultHeadersMiddleware = func(c *resty.Client, request *resty.Request) error {
		request.SetHeaders(map[string]string{
			"X-Time":       strconv.FormatInt(time.Now().UnixMilli(), 10),
			"X-Nonce":      uuid.NewString(),
			"X-Request-Id": uuid.NewString(),
		})

		return nil
	}
	var signRequestMiddleware = func(c *resty.Client, request *resty.Request) error {
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

	c := resty.New()

	c.OnBeforeRequest(defaultHeadersMiddleware).
		OnBeforeRequest(signRequestMiddleware)

	return &client{
		credentials: &Credentials{},
		Mining: &mining{
			httpClient: c,
		},
	}
}
