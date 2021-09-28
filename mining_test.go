package nhclient

import (
	"fmt"
	"testing"
)

func Test_mining_GetAddress(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, err := New().Authenticate(&Credentials{
				orgId:     "1ea33add-2ea3-419a-8574-e0b6efb5b125",
				apiKey:    "6b442ff4-4512-4de8-8e58-1d42f6549567",
				apiSecret: "ceb7c515-c5a7-41e7-a354-b2e293dd6d7fc647ba7c-fca6-492f-9bb0-8fa6d0fabe2a",
			}).Mining.GetAddress()

			fmt.Println(a, err)
		})
	}
}
