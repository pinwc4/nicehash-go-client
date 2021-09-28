package nhclient

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type public struct {
	httpClient *resty.Client
}

func (p *public) GetMiningAlgorithms() {
	resp, err := p.httpClient.R().Get("https://api2.nicehash.com/main/api/v2/mining/algorithms")

	fmt.Println(resp, err)
}
