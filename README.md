# NiceHash Go Client (In development)

TABLE OF CONTENTS
---------------------

- Introduction
- Requirements
- Installation
- Configuration
- Maintainers
- TODO Endpoints

INTRODUCTION
---------------------

This client allows you easily integrate the [NiceHash API](https://www.nicehash.com/docs/) into your GoLang application.

REQUIREMENTS
---------------------

- GoLang 1.17

INSTALLATION
---------------------

```bash
 go get github.com/GutoScherer/nicehash-client
```

CONFIGURATION
---------------------

```go
package main

import "github.com/GutoScherer/nicehash-client"

func main() {
	client := nhclient.New()

	//With this client you can access the public endpoints
	algos, err := client.Public.Mining.GetAlgorithms()
	
	//To access the private endpoints you simply need to call the Authenticate method with your credentials
	client.Authenticate("ORG_ID", "API_KEY", "SECRET_KEY")
	
	miningAddress, err := client.Private.Mining.GetAddress()
}
```

MAINTAINERS
---------------------

- Luiz Augusto Solheid Scherer - https://github.com/GutoScherer

TODO
----

### Accounting

- [ ] GET /main/api/v2/accounting/account2/{currency}
- [ ] GET /main/api/v2/accounting/accounts2
- [ ] GET /main/api/v2/accounting/activity/{currency}
- [ ] GET /main/api/v2/accounting/depositAddresses
- [ ] GET /main/api/v2/accounting/deposits/{currency}
- [ ] GET /main/api/v2/accounting/deposits2/{currency}/{id}
- [ ] GET /main/api/v2/accounting/exchange/{id}/trades
- [ ] GET /main/api/v2/accounting/hashpower/{id}/transactions
- [ ] GET /main/api/v2/accounting/hashpowerEarnings/{currency}
- [ ] GET /main/api/v2/accounting/transaction/{currency}/{transactionId}
- [ ] GET /main/api/v2/accounting/transactions/{currency}
- [ ] POST /main/api/v2/accounting/withdrawal
- [ ] DELETE /main/api/v2/accounting/withdrawal/{currency}/{id}
- [ ] GET /main/api/v2/accounting/withdrawal2/{currency}/{id}
- [ ] GET /main/api/v2/accounting/withdrawalAddress/{id}
- [ ] GET /main/api/v2/accounting/withdrawalAddresses
- [ ] GET /main/api/v2/accounting/withdrawals/{currency}

### Miner private

- [ ] GET /main/api/v2/mining/miningAddress
- [ ] GET /main/api/v2/mining/algo/stats
- [ ] GET /main/api/v2/mining/groups/list
- [ ] GET /main/api/v2/mining/rig/stats/algo
- [ ] GET /main/api/v2/mining/rig/stats/unpaid
- [ ] GET /main/api/v2/mining/rig2/{rigId}
- [ ] GET /main/api/v2/mining/rigs/activeWorkers
- [ ] GET /main/api/v2/mining/rigs/payouts
- [ ] GET /main/api/v2/mining/rigs/stats/algo
- [ ] GET /main/api/v2/mining/rigs/stats/unpaid
- [ ] POST /main/api/v2/mining/rigs/status2
- [ ] GET /main/api/v2/mining/rigs2

### Pools

- [ ] POST /main/api/v2/pool
- [ ] GET /main/api/v2/pool/{poolId}
- [ ] DELETE /main/api/v2/pool/{poolId}
- [ ] GET /main/api/v2/pools
- [ ] POST /main/api/v2/pools/verify
- [ ] POST /main/api/v2/pool
- [ ] GET /main/api/v2/pool/{poolId}
- [ ] DELETE /main/api/v2/pool/{poolId}
- [ ] GET /main/api/v2/pools
- [ ] POST /main/api/v2/pools/verify

### Public

- [ ] GET /main/api/v2/mining/algorithms
- [ ] GET /main/api/v2/mining/markets
- [ ] GET /main/api/v2/public/currencies
- [ ] GET /main/api/v2/public/service/fee/info
- [ ] GET /api/v2/enum/countries
- [ ] GET /api/v2/enum/kmCountries
- [ ] GET /api/v2/enum/permissions
- [ ] GET /api/v2/enum/xchCountries
- [ ] GET /api/v2/system/flags
- [ ] GET /api/v2/time

### Reports

- [ ] POST /main/api/v2/reports/add
- [ ] DELETE /main/api/v2/reports/delete/{id}
- [ ] GET /main/api/v2/reports/download/pdf/{id}
- [ ] GET /main/api/v2/reports/download/{id}
- [ ] GET /main/api/v2/reports/list

### External miner

- [ ] GET /main/api/v2/mining/external/{btcAddress}/rigs/activeWorkers
- [ ] GET /main/api/v2/mining/external/{btcAddress}/rigs/stats/algo
- [ ] GET /main/api/v2/mining/external/{btcAddress}/rigs/stats/unpaid
- [ ] GET /main/api/v2/mining/external/{btcAddress}/rigs/withdrawals
- [ ] GET /main/api/v2/mining/external/{btcAddress}/rigs2

### Hashpower private

- [ ] GET /main/api/v2/hashpower/myOrders
- [ ] POST /main/api/v2/hashpower/order
- [ ] GET /main/api/v2/hashpower/order/{id}
- [ ] DELETE /main/api/v2/hashpower/order/{id}
- [ ] POST /main/api/v2/hashpower/order/{id}/refill
- [ ] GET /main/api/v2/hashpower/order/{id}/stats
- [ ] POST /main/api/v2/hashpower/order/{id}/updatePriceAndLimit
- [ ] POST /main/api/v2/hashpower/orders/calculateEstimateDuration

### Hashpower public

- [ ] GET /main/api/v2/hashpower/orderBook
- [ ] POST /main/api/v2/hashpower/orders/fixedPrice
- [ ] GET /main/api/v2/hashpower/orders/summaries
- [ ] GET /main/api/v2/hashpower/orders/summary
- [ ] GET /main/api/v2/public/algo/history
- [ ] GET /main/api/v2/public/buy/info
- [ ] GET /main/api/v2/public/orders
- [ ] GET /main/api/v2/public/simplemultialgo/info
- [ ] GET /main/api/v2/public/stats/global/24h
- [ ] GET /main/api/v2/public/stats/global/current

### Exchange private

- [ ] DELETE /exchange/api/v2/info/cancelAllOrders
- [ ] GET /exchange/api/v2/info/fees/status
- [ ] GET /exchange/api/v2/info/myOrder
- [ ] GET /exchange/api/v2/info/myOrders
- [ ] GET /exchange/api/v2/info/myTrades
- [ ] GET /exchange/api/v2/info/orderTrades
- [ ] POST /exchange/api/v2/order
- [ ] DELETE /exchange/api/v2/order

### Exchange public

- [ ] GET /exchange/api/v2/info/candlesticks
- [ ] GET /exchange/api/v2/info/marketStats
- [ ] GET /exchange/api/v2/info/prices
- [ ] GET /exchange/api/v2/info/status
- [ ] GET /exchange/api/v2/info/trades
- [ ] GET /exchange/api/v2/orderbook
