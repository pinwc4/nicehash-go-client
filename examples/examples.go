package main

import (
	"fmt"
	"github.com/GutoScherer/nicehash-client"
	"github.com/davecgh/go-spew/spew"
	"os"
)

func main() {
	client := nhclient.New().Authenticate(
		os.Getenv("ORG_ID"),
		os.Getenv("API_KEY"),
		os.Getenv("API_SECRET"),
	)

	rig, err := client.Private.Mining.GetRigDetails(os.Getenv("RIG_ID"))
	if err != nil {
		fmt.Println(err)
	} else {
		spew.Dump(rig)
	}
}
