package nhclient

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/pkg/errors"
)

type hashpowerPrivate struct {
	client *client
}

type newOrderInfo struct {
	Id string
}

func (h *hashpowerPrivate) CreateOrder(algo *miningAlgoInfo, market string, poolid string, price float64, limit float64, btcamount float64) (orderInfo *newOrderInfo, err error) {
	path := "/main/api/v2/hashpower/order"

	requestbodymap := map[string]interface{}{
		"market":              market,
		"algorithm":           algo.Algorithm,
		"amount":              fmt.Sprintf("%.3f", btcamount),
		"displayMarketFactor": algo.DisplayMarketFactor,
		"marketFactor":        algo.MarketFactor,
		"price":               fmt.Sprintf("%.4f", price),
		"poolId":              poolid,
		"limit":               fmt.Sprintf("%.3f", limit),
		"type":                "STANDARD",
	}

	requestbody, _ := json.Marshal(requestbodymap)

	responseBody, err := h.client.doRequest(
		"POST",
		path,
		requestbody,
		nil,
	)

	if err != nil {
		return nil, errors.Wrap(err, "executing request")
	}

	err = json.Unmarshal(responseBody, &orderInfo)

	if err != nil {
		return nil, errors.Wrap(err, "unmarshalling response body")
	}

	return orderInfo, nil
}
