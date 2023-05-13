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
		"org",
		"key",
		"secret",
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

	fmt.Println(os.Getenv("$TESTENV"))
	fmt.Println("---------------")

	info5, err5 := clientp.Private.HashPower.GetMyOrders("ETCHASH", "EU")

	if err5 != nil {
		fmt.Println("error gettin")
		fmt.Println(err5)
	} else {
		fmt.Println("gotten")
		fmt.Println(info5)
		fmt.Println(len(info5.List))
		//fmt.Println(info5.List[1])
	}
	fmt.Println("---------------")
}
