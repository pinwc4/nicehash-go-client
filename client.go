package nhclient

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

const ProdUrl = "https://api2.nicehash.com"
const TestUrl = "https://api-test.nicehash.com"

type client struct {
	httpClient *resty.Client
	Private    *struct {
		Mining    *miningPrivate
		Exchange  *exchangePrivate
		HashPower *hashpowerPrivate
		Reports   *reports
		Pools     *pools
	}
	Public *struct {
		General   *general
		Mining    *miningPublic
		Exchange  *exchangePublic
		HashPower *hashpowerPublic
	}
}

func New() *client {
	client := &client{
		httpClient: resty.New().
			OnBeforeRequest(defaultHeaders).
			SetError(&requestError{}),
	}

	client.Private = &struct {
		Mining    *miningPrivate
		Exchange  *exchangePrivate
		HashPower *hashpowerPrivate
		Reports   *reports
		Pools     *pools
	}{
		Mining:    &miningPrivate{client: client},
		Exchange:  &exchangePrivate{client: client},
		HashPower: &hashpowerPrivate{client: client},
		Reports:   &reports{client: client},
		Pools:     &pools{client: client},
	}

	client.Public = &struct {
		General   *general
		Mining    *miningPublic
		Exchange  *exchangePublic
		HashPower *hashpowerPublic
	}{
		General:   &general{client: client},
		Mining:    &miningPublic{client: client},
		Exchange:  &exchangePublic{client: client},
		HashPower: &hashpowerPublic{client: client},
	}

	return client
}

func (c *client) Authenticate(orgId, apiKey, secretKey string) *client {
	c.httpClient.OnBeforeRequest(authenticator(orgId, apiKey, secretKey))

	return c
}

func (c *client) doRequest(method, path string, body interface{}, queryParameters map[string]string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", ProdUrl, path)

	response, err := c.httpClient.R().
		SetQueryParams(queryParameters).
		SetBody(body).
		Execute(method, url)

	if err != nil {
		return nil, errors.Wrap(err, "error executing http request")
	}

	if err, ok := response.Error().(*requestError); ok {
		return nil, errors.Wrap(err, "API error")
	}

	return response.Body(), nil
}
