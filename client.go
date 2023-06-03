package nhclient

import (
	"fmt"
	"net"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

const ProdUrl = "https://api2.nicehash.com"
const TestUrl = "https://api-test.nicehash.com"

type Client struct {
	httpClient *resty.Client
	Private    *struct {
		Accounting *accounting
		Mining     *miningPrivate
		Exchange   *exchangePrivate
		HashPower  *hashpowerPrivate
		Reports    *reports
		Pools      *pools
	}
	Public *struct {
		General   *general
		Mining    *miningPublic
		Exchange  *exchangePublic
		HashPower *hashpowerPublic
	}
}

func New() *Client {
	return NewWithAddress("")
}

func NewWithAddress(connectaddress string) *Client {
	client := &Client{
		httpClient: resty.New().
			OnBeforeRequest(defaultHeaders).
			SetError(&requestError{}),
	}
	if connectaddress != "" {
		tempaddr := connectaddress + ":0"
		addr, err := net.ResolveTCPAddr("tcp", tempaddr)
		if err != nil {
			fmt.Println("error using", connectaddress)
			return nil
		}
		dialer := &net.Dialer{LocalAddr: addr}
		ptransport := &http.Transport{Dial: dialer.Dial}
		client.httpClient.SetTransport(ptransport)

	}

	client.Private = &struct {
		Accounting *accounting
		Mining     *miningPrivate
		Exchange   *exchangePrivate
		HashPower  *hashpowerPrivate
		Reports    *reports
		Pools      *pools
	}{
		Accounting: &accounting{client: client},
		Mining:     &miningPrivate{client: client},
		Exchange:   &exchangePrivate{client: client},
		HashPower:  &hashpowerPrivate{client: client},
		Reports:    &reports{client: client},
		Pools:      &pools{client: client},
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

func (c *Client) Authenticate(orgId, apiKey, secretKey string) *Client {
	c.httpClient.OnBeforeRequest(authenticator(orgId, apiKey, secretKey))

	return c
}

func (c *Client) doRequest(method, path string, body interface{}, queryParameters map[string]string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", ProdUrl, path)

	response, err := c.httpClient.R().
		SetQueryParams(queryParameters).
		SetBody(body).
		Execute(method, url)

	if err != nil {
		return nil, errors.Wrap(err, "error executing http request")
	}

	if response.StatusCode() == http.StatusNotFound {
		return nil, errInvalidAPIPath
	}

	if err, ok := response.Error().(*requestError); ok {
		return nil, errors.Wrap(err, "API error")
	}

	return response.Body(), nil
}

func WithOptionalParameter(key, value string) *optionalParameter {
	return &optionalParameter{
		key:   key,
		value: value,
	}
}
