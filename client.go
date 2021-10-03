package nhclient

import (
	"github.com/GutoScherer/nicehash-client/middlewares"
	"github.com/go-resty/resty/v2"
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
			OnBeforeRequest(middlewares.NewDefaultHeaders()).
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

func (c *client) Authenticate(orgId, apiKey, secretKey string) {
	c.httpClient.OnBeforeRequest(middlewares.NewAuthenticator(nil))
}

func (c *client) doRequest() {

}
