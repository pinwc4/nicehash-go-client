package nhclient

import (
	//"encoding/json"
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
)

type general struct {
	client *Client
}

type miningAlgoList struct {
	MiningAlgorithms []*miningAlgoInfo
}

type miningAlgoInfo struct {
	Algorithm             string
	Title                 string
	Enabled               bool
	Order                 int64
	DisplayMiningFactor   string `json:"displayMiningFactor"`
	MiningFactor          string `json:"miningFactor"`
	DisplayMarketFactor   string `json:"displayMarketFactor"`
	MarketFactor          string `json:"marketFactor"`
	MinimalOrderAmount    string `json:"minimalOrderAmount"`
	MinSpeedLimit         string `json:"minSpeedLimit"`
	MaxSpeedLimit         string `json:"maxSpeedLimit"`
	PriceDownStep         string `json:"priceDownStep"`
	MinimalPoolDifficulty string `json:"minimalPoolDifficulty"`
	Port                  int64
	Color                 string
	OrdersEnabled         bool
	EnabledMarkets        string `json:"enabledMarkets"`
	DisplayPriceFactor    string `json:"displayPriceFactor"`
	PriceFactor           string `json:"priceFactor"`
}

func (g *general) GetMiningAlgorithms() (algomap map[string]*miningAlgoInfo, err error) {
	algomap = make(map[string]*miningAlgoInfo)
	responseBody, err := g.client.doRequest(
		"GET",
		"/main/api/v2/mining/algorithms",
		nil,
		nil,
	)

	if err != nil {
		return nil, errors.Wrap(err, "executing request")
	}

	var algos miningAlgoList

	err = json.Unmarshal(responseBody, &algos)

	if err != nil {
		return nil, errors.Wrap(err, "unmarshalling response body")
	}

	for _, algo := range algos.MiningAlgorithms {
		algomap[algo.Algorithm] = algo
	}
	return algomap, nil

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
