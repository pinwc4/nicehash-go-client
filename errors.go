package nhclient

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var errNotImplemented = errors.New("not implemented yet")
var errInvalidAPIPath = errors.New("invalid API path")

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


