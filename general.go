package nhclient

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type general struct {
	client *client
}

type miningAlgoList struct {
	MiningAlgorithms []*miningAlgoInfo
}

type miningAlgoInfo struct {
	Algorithm             string
	Title                 string
	Enabled               bool
	Order                 int64
	DisplayMiningFactor   string
	MiningFactor          float64
	DisplayMarketFactor   string
	MarketFactor          float64
	MinimalOrderAmount    float64
	MinSpeedLimit         float64
	MaxSpeedLimit         float64
	PriceDownStep         float64
	MinimalPoolDifficulty float64
	Port                  int64
	Color                 string
	OrdersEnabled         bool
	EnabledMarkets        string
	DisplayPriceFactor    string
	PriceFactor           float64
}

func (g *general) GetMiningAlgorithms() (algos *miningAlgoList, err error) {
	responseBody, err := g.client.doRequest(
		"GET",
		"/main/api/v2/mining/algorithms",
		nil,
		nil,
	)

	if err != nil {
		return nil, errors.Wrap(err, "executing request")
	}

	err = json.Unmarshal(responseBody, &algos)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshalling response body")
	}

	return algos, nil

}

func (g *general) GetMiningMarkets() error {
	return errNotImplemented
}

func (g *general) GetCurrencies() error {
	return errNotImplemented
}
func (g *general) GetFeeRules() error {
	return errNotImplemented
}
func (g *general) GetCountries() error {
	return errNotImplemented
}
func (g *general) GetAllowedKMCountries() error {
	return errNotImplemented
}
func (g *general) GetPossiblePermissions() error {
	return errNotImplemented
}
func (g *general) GetAllowedXCHCountries() error {
	return errNotImplemented
}
func (g *general) GetAPIFlags() error {
	return errNotImplemented
}
func (g *general) GetServerTime() error {
	return errNotImplemented
}
