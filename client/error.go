package client

import (
	"fmt"
	"net/http"
)

type HTTPError struct {
	status     string
	StatusCode int
	url        string
}

func httpError(req *http.Request, resp *http.Response) error {
	return &HTTPError{
		status:     resp.Status,
		StatusCode: resp.StatusCode,
		url:        req.URL.String(),
	}
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%s: %s", e.url, e.status)
}
