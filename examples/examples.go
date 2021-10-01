package main

import (
	"fmt"
	"github.com/GutoScherer/nicehash-client"
	"os"
)

func main() {
	nhclient.Authenticate(&nhclient.Credentials{
		OrgId:     os.Getenv("ORG_ID"),
		ApiKey:    os.Getenv("API_KEY"),
		ApiSecret: os.Getenv("API_SECRET"),
	})

	address, err := nhclient.Mining.GetAddress()

	fmt.Println(address, err)
}
