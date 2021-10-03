package middlewares

import (
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"strconv"
	"time"
)

func NewDefaultHeaders() resty.RequestMiddleware {
	return func(c *resty.Client, request *resty.Request) error {
		request.SetHeaders(map[string]string{
			"X-Time":       strconv.FormatInt(time.Now().UnixMilli(), 10),
			"X-Nonce":      uuid.NewString(),
			"X-Request-Id": uuid.NewString(),
		})

		return nil
	}
}
