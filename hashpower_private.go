package nhclient

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/pkg/errors"
)

type hashpowerPrivate struct {
	client *Client
}

type newOrderInfo struct {
	Id string
}

type myOrders struct {
	List []*struct {
		Id                   string
		AvailableAmount      string
		PayedAmount          string
		Price                string
		Limit                string
		Amount               string
		AcceptedCurrentSpeed string
		Pool                 struct {
			Id        string
			Name      string
			Algorithm string
		}
		Alive     bool
		RigsCount int64
	}
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

func (h *hashpowerPrivate) CancelOrder(orderid string) (err error) {
	path := fmt.Sprintf("/main/api/v2/hashpower/order/%s", orderid)

	_, err = h.client.doRequest(
		"DELETE",
		path,
		nil,
		nil,
	)

	if err != nil {
		return errors.Wrap(err, "executing request")
	}
	return nil
}

func (h *hashpowerPrivate) UpdateOrderPrice(orderid string, price float64, algo *miningAlgoInfo) (err error) {
	path := fmt.Sprintf("/main/api/v2/hashpower/order/%s/updatePriceAndLimit/", orderid)

	requestbody := fmt.Sprintf(
		`{"price":"%s","marketFactor":"%s","displayMarketFactor":"%s"}`,
		fmt.Sprintf("%.4f", price),
		algo.MarketFactor,
		algo.DisplayMarketFactor,
	)

	_, err = h.client.doRequest(
		"POST",
		path,
		requestbody,
		nil,
	)

	if err != nil {
		return errors.Wrap(err, "executing request")
	}
	return nil

}

func (h *hashpowerPrivate) UpdateOrderLimit(orderid string, limit float64, algo *miningAlgoInfo) (err error) {
	path := fmt.Sprintf("/main/api/v2/hashpower/order/%s/updatePriceAndLimit/", orderid)

	requestbody := fmt.Sprintf(
		`{"limit":"%s","marketFactor":"%s","displayMarketFactor":"%s"}`,
		fmt.Sprintf("%.4f", limit),
		algo.MarketFactor,
		algo.DisplayMarketFactor,
	)

	_, err = h.client.doRequest(
		"POST",
		path,
		requestbody,
		nil,
	)

	if err != nil {
		return errors.Wrap(err, "executing request")
	}
	return nil

}

func (h *hashpowerPrivate) UpdateOrderPriceAndLimit(orderid string, price float64, limit float64, algo *miningAlgoInfo) (err error) {
	path := fmt.Sprintf("/main/api/v2/hashpower/order/%s/updatePriceAndLimit/", orderid)

	requestbody := fmt.Sprintf(
		`{"price":"%s","limit":"%s","marketFactor":"%s","displayMarketFactor":"%s"}`,
		fmt.Sprintf("%.4f", price),
		fmt.Sprintf("%.4f", limit),
		algo.MarketFactor,
		algo.DisplayMarketFactor,
	)

	fmt.Println(requestbody)

	_, err = h.client.doRequest(
		"POST",
		path,
		requestbody,
		nil,
	)

	if err != nil {
		return errors.Wrap(err, "executing request")
	}
	return nil

}

func (h *hashpowerPrivate) GetMyOrders(algo string, market string) (myOrders *myOrders, err error) {
	path := "/main/api/v2/hashpower/myOrders"

	queryparameters := make(map[string]string, 0)
	queryparameters["algorithm"] = algo
	queryparameters["market"] = market
	queryparameters["ts"] = "0"
	queryparameters["op"] = "GT"
	queryparameters["limit"] = "1000"
	queryparameters["active"] = "true"

	responseBody, err := h.client.doRequest(
		"GET",
		path,
		nil,
		queryparameters,
	)

	if err != nil {
		return nil, errors.Wrap(err, "executing request")
	}

	err = json.Unmarshal(responseBody, &myOrders)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshalling response body")
	}

	return myOrders, nil
}
