package nhclient

type requestError struct {
	ErrorID string `json:"error_id"`
	Errors  []*struct {
		Code    int64
		Message string
	}
}
