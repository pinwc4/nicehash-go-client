package main

import (
	"fmt"
	"os"

	nhclient "github.com/pinwc4/nicehash-go-client"
)

var client = nhclient.New()

func main() {

	clientp := nhclient.New()
	clientp.Authenticate(
		"orgid",
		"apikey",
		"secretkey",
	)

	miningAddress, err := clientp.Private.Mining.GetAddress()
	fmt.Println("---------------")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(miningAddress)
	}

	//rig, err := client.Private.Mining.GetRigDetails(os.Getenv("RIG_ID"))
	info2, err2 := client.Public.General.GetMiningAlgorithms()
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("---------------")
		fmt.Println(info2)

		fmt.Println("---------------")
		fmt.Println(info2["ETCHASH"])
		fmt.Println(info2["ETCHASH"].Algorithm)
		fmt.Println(info2["ETCHASH"].MarketFactor)
		fmt.Println("---------------")
	}

	info3, err3 := client.Public.HashPower.GetOrderBook("ETCHASH")
	if err3 != nil {
		fmt.Println(err)
	}

	fmt.Println(info3)
	fmt.Println(info3.Stats.EU.Orders[0])
	fmt.Println("---------------")

	info4, err4 := clientp.Private.HashPower.CreateOrder(info2["ETCHASH"], "EU", "poolid", 1.0011, 1.02, 1.001)

	if err4 != nil {
		fmt.Println("error with creating order")
		fmt.Println(err4)
	}

	fmt.Println(info4)
	fmt.Println("---------------")

	fmt.Println(os.Getenv("$TESTENV"))
}
