package main

import (
	"fmt"
	"os"

	nhclient "github.com/pinwc4/nicehash-client"
)

var client = nhclient.New().Authenticate(
	os.Getenv("ORG_ID"),
	os.Getenv("API_KEY"),
	os.Getenv("API_SECRET"),
)

func main() {
	rig, err := client.Private.Mining.GetRigDetails(os.Getenv("RIG_ID"))
	if err != nil {
		fmt.Println(err)
	} else {
		for _, device := range rig.Devices {
			fmt.Println(device.Name + " is " + device.Status.Description)
		}
		//spew.Dump(rig)
	}

	fmt.Println("---------------")

	client.Private.Mining.GetRigsGroups(nhclient.WithOptionalParameter("extendedResponse", "true"))
}
