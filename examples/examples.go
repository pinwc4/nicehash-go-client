package main

import (
	"fmt"
	//"os"

	nhclient "github.com/pinwc4/nicehash-go-client"
)

//var client = nhclient.New().Authenticate(
//	os.Getenv("ORG_ID"),
//	os.Getenv("API_KEY"),
//	os.Getenv("API_SECRET"),
//)

var client = nhclient.New()

func main() {
	//rig, err := client.Private.Mining.GetRigDetails(os.Getenv("RIG_ID"))
	info2, err := client.Public.General.GetMiningAlgorithms()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("---------------")
		fmt.Println(info2)

		fmt.Println("---------------")
		fmt.Println(info2["SCRYPT"])
		fmt.Println(info2["SCRYPT"].Algorithm)
		fmt.Println(info2["SCRYPT"].MarketFactor)
		fmt.Println("---------------")
	}

	info3, err3 := client.Public.HashPower.GetOrderBook("DAGGERHASHIMOTO")
	if err3 != nil {
		fmt.Println(err)
	}

	fmt.Println(info3)
	fmt.Println(info3.Stats.EU.Orders[0])

	client.Private.Mining.GetRigsGroups(nhclient.WithOptionalParameter("extendedResponse", "true"))
}
