package nhclient

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type accounting struct {
	httpClient *resty.Client
}

func (a *accounting) GetTotalBalance() {
	resp, err := a.httpClient.R().Get("https://api2.nicehash.com/main/api/v2/accounting/accounts2")
	fmt.Println(resp, err)
}
