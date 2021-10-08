package main

import (
	"fmt"
	"github.com/GutoScherer/nicehash-client"
	"github.com/davecgh/go-spew/spew"
	"os"
)

var client = nhclient.New().Authenticate(
	os.Getenv("ORG_ID"),
	os.Getenv("API_KEY"),
	os.Getenv("API_SECRET"),
)

func main() {
	printRigDetails()
	/*printMiningAddress()
	printAlgos()*/
}

func printAlgos() {
	client.Private.Mining.GetAlgos("")
}

func printMiningAddress() {
	addr, err := client.Private.Mining.GetAddress()
	if err != nil {
		fmt.Println(err)
	} else {
		spew.Dump(addr)
	}
}

func printRigDetails() {
	rig, err := client.Private.Mining.GetRigDetails(os.Getenv("RIG_ID"))
	if err != nil {
		fmt.Println(err)
	} else {
		spew.Dump(rig)
	}
}
