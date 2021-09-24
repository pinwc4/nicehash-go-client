package nicehash_client

import "github.com/go-resty/resty/v2"

type Credentials struct {
	orgId     string
	apiKey    string
	apiSecret string
}

type client struct {
	httpClient *resty.Client

	credentials *Credentials

	mining *mining
}

func New() *client {
	return &client{
		credentials: &Credentials{},
		mining:      &mining{},
	}
}

func (c *client) Mining() *mining {
	return c.mining
}

func (c *client) signed() *resty.Request {
	return c.httpClient.R().SetHeaders(
		map[string]string{
			"X-Nonce":           "",
			"X-Organization-Id": "",
			"X-Auth":            "",
		},
	)
}
