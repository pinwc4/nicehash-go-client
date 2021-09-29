package nhclient

import (
	"fmt"
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	client := New().Authenticate(&Credentials{
		orgId:     os.Getenv("ORG_ID"),
		apiKey:    os.Getenv("API_KEY"),
		apiSecret: os.Getenv("API_SECRET"),
	})

	rig, err := client.Mining.GetRigDetails("0-sDiZZc7S-Eq7jgtZ47qN3A")

	fmt.Println(rig.Devices[1].GetCoreTemperature(), rig.Devices[1].GetMemoryTemperature(), err)
}
