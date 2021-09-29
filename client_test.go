package nhclient

import (
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	New().Authenticate(&Credentials{
		orgId:     os.Getenv("ORG_ID"),
		apiKey:    os.Getenv("API_KEY"),
		apiSecret: os.Getenv("API_SECRET"),
	}).Accounting.GetTotalBalance()
}
