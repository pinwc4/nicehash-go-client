package nhclient

import "encoding/json"

type requestError struct {
	ErrorID string `json:"error_id"`
	Errors  []*struct {
		Code    int64
		Message string
	}
}

func (r *requestError) Error() string {
	errors, _ := json.Marshal(r.Errors)

	return string(errors)
}
