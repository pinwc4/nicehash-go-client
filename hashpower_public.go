package nhclient

import (
	//"encoding/json"
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
)

type hashpowerPublic struct {
	client *client
}

type orderBook struct {
	Stats *struct {
		EU  *orderBookDetails
		USA *orderBookDetails
	}
}

type orderBookDetails struct {
	UpdatedTs  string
	TotalSpeed string
	//MarketFactor        string
	//DisplayMarketFactor string
	//PriceFactor         string
	//DisplayPriceFactor  string
	Orders []*struct {
		Id            string
		Type          string
		Price         string
		Limit         string
		RigsCount     int64
		AcceptedSpeed string
		PayingSpeed   string
		Alive         bool
	}
	//Pagination *struct {
	//	size           int64
	//	page           int64
	//	totalPageCount int64
	//}
}

func (h *hashpowerPublic) GetOrderBook(algo string) (orderBook *orderBook, err error) {
	path := "/main/api/v2/hashpower/orderBook"

	queryParameters := make(map[string]string, 0)
	queryParameters["algorithm"] = algo
	queryParameters["size"] = "1000"

	responseBody, err := h.client.doRequest(
		"GET",
		path,
		nil,
		queryParameters,
	)

	if err != nil {
		return nil, errors.Wrap(err, "executing request")
	}

	err = json.Unmarshal(responseBody, &orderBook)

	if err != nil {
		return nil, errors.Wrap(err, "unmarshalling response body")
	}

	return orderBook, nil
}
