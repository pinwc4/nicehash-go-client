package main

import (
	"fmt"
	"github.com/GutoScherer/nicehash-client"
	"os"
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
			if device.Status.EnumName == nhclient.DeviceStatusMining {
				fmt.Println(device.Name + " is mining!")
			} else {
				fmt.Println(device.Name + " is " + device.Status.Description)
			}
		}
		//spew.Dump(rig)
	}
}