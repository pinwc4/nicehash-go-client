package nhclient

import (
	"fmt"
)

type public struct{}

var Public public

func (p *public) GetMiningAlgorithms() {
	resp, err := httpClient.R().Get("https://api2.nicehash.com/main/api/v2/mining/algorithms")

	fmt.Println(resp, err)
}
