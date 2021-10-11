package nhclient

import "testing"

func Test_general_GetMiningAlgorithms(t *testing.T) {
	type fields struct {
		client *client
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "Test 1",
			fields: fields{
				client: New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &general{
				client: tt.fields.client,
			}

			g.GetMiningAlgorithms()
		})
	}
}
