package nhclient

import (
	"fmt"
)

type accounting struct{}

var Accounting accounting

func (a *accounting) GetTotalBalance() {
	resp, err := httpClient.R().Get("https://api2.nicehash.com/main/api/v2/accounting/accounts2")
	fmt.Println(resp, err)
}
