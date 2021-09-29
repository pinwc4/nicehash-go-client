package nhclient

import (
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type Credentials struct {
	orgId     string
	apiKey    string
	apiSecret string
}

type client struct {
	httpClient *resty.Client
	Mining     *mining
	Public     *public
	Accounting *accounting
}

func New() *client {
	httpClient := resty.New().OnBeforeRequest(func(c *resty.Client, request *resty.Request) error {
		request.SetHeaders(map[string]string{
			"X-Time":       strconv.FormatInt(time.Now().UnixMilli(), 10),
			"X-Nonce":      uuid.NewString(),
			"X-Request-Id": uuid.NewString(),
		})

		return nil
	})

	return &client{
		httpClient: httpClient,
		Mining: &mining{
			httpClient: httpClient,
		},
		Public: &public{
			httpClient: httpClient,
		},
		Accounting: &accounting{
			httpClient: httpClient,
		},
	}
}

func (c *client) Authenticate(credentials *Credentials) *client {
	c.httpClient.OnBeforeRequest(authenticatorMiddleware(credentials))

	return c
}
