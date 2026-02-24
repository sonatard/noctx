package httptest_request

import (
	"context"
	"net/http"
	"net/http/httptest"
)

func _() {
	const url = "https://example.com"

	cli := &http.Client{}

	ctx := context.Background()

	req := httptest.NewRequest(http.MethodPost, url, nil) // want `net/http/httptest\.NewRequest must not be called. use net/http/httptest\.NewRequestWithContext`
	cli.Do(req)

	req2 := httptest.NewRequestWithContext(ctx, http.MethodPost, url, nil) // OK
	cli.Do(req2)

	newRequest := httptest.NewRequest
	req3 := newRequest(http.MethodPost, url, nil) // want `net/http/httptest\.NewRequest must not be called. use net/http/httptest\.NewRequestWithContext`
	cli.Do(req3)
}
