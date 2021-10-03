package main

import "github.com/GutoScherer/nicehash-client"

func main() {
	client := nhclient.New()

	client.Public.General.GetMiningAlgorithms()
}
