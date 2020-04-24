package a

import (
	"context"
	"net/http"
)

func main() {
	const url = "http://example.com"
	cli := &http.Client{}

	ctx := context.Background()
	http.Get(url) // want `net/http\.Get must not be called`
	_ = http.Get  // OK
	f := http.Get // OK
	f(url)        // want `net/http\.Get must not be called`

	http.Head(url)          // want `net/http\.Head must not be called`
	http.Post(url, "", nil) // want `net/http\.Post must not be called`
	http.PostForm(url, nil) // want `net/http\.PostForm must not be called`

	cli.Get(url) // want `\(\*net/http\.Client\)\.Get must not be called`
	_ = cli.Get  // OK
	m := cli.Get // OK
	m(url)       // want `\(\*net/http\.Client\)\.Get must not be called`

	cli.Head(url)          // want `\(\*net/http\.Client\)\.Head must not be called`
	cli.Post(url, "", nil) // want `\(\*net/http\.Client\)\.Post must not be called`
	cli.PostForm(url, nil) // want `\(\*net/http\.Client\)\.PostForm must not be called`

	req, _ := http.NewRequest(http.MethodPost, url, nil) // want `should rewrite http.NewRequest to http.NewRequestWithContext or http.NewRequest and \(\*Request\).WithContext`
	cli.Do(req)

	req2, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, nil) // OK
	cli.Do(req2)

	req3, _ := http.NewRequest(http.MethodPost, url, nil) // OK
	req3 = req3.WithContext(ctx)
	cli.Do(req3)

	f2 := func(req *http.Request, ctx context.Context) *http.Request {
		return req
	}
	req4, _ := http.NewRequest(http.MethodPost, url, nil) // want `should rewrite http.NewRequest to http.NewRequestWithContext or http.NewRequest and \(\*Request\).WithContext`
	req4 = f2(req4, ctx)

	newRequest := http.NewRequest
	req5, _ := newRequest(http.MethodPost, url, nil) // want `should rewrite http.NewRequest to http.NewRequestWithContext or http.NewRequest and \(\*Request\).WithContext`
	cli.Do(req5)

	type MyRequest = http.Request
	f3 := func(req *MyRequest, ctx context.Context) *MyRequest {
		return req
	}
	req6, _ := http.NewRequest(http.MethodPost, url, nil) // want `should rewrite http.NewRequest to http.NewRequestWithContext or http.NewRequest and \(\*Request\).WithContext`
	req6 = f3(req6, ctx)

	type MyRequest2 http.Request
	f4 := func(req *MyRequest2, ctx context.Context) *MyRequest2 {
		return req
	}
	req7, _ := http.NewRequest(http.MethodPost, url, nil) // want `should rewrite http.NewRequest to http.NewRequestWithContext or http.NewRequest and \(\*Request\).WithContext`
	req72 := MyRequest2(*req7)
	f4(&req72, ctx)

	req8, _ := func() (*http.Request, error) {
		return http.NewRequest(http.MethodPost, url, nil) // want `should rewrite http.NewRequest to http.NewRequestWithContext or http.NewRequest and \(\*Request\).WithContext`
	}()
	cli.Do(req8)

	f5 := func(req, req2 *http.Request, ctx context.Context) (*http.Request, *http.Request) {
		return req, req2
	}
	req9, _ := http.NewRequest(http.MethodPost, url, nil) // want `should rewrite http.NewRequest to http.NewRequestWithContext or http.NewRequest and \(\*Request\).WithContext`
	req9, _ = f5(req9, req9, ctx)

	req10, _ := http.NewRequest(http.MethodPost, url, nil) // want `should rewrite http.NewRequest to http.NewRequestWithContext or http.NewRequest and \(\*Request\).WithContext`
	req11, _ := http.NewRequest(http.MethodPost, url, nil) // want `should rewrite http.NewRequest to http.NewRequestWithContext or http.NewRequest and \(\*Request\).WithContext`
	req9, req11 = f5(req10, req11, ctx)

	req14, req15 := func() (*http.Request, *http.Request) {
		req12, _ := http.NewRequest(http.MethodPost, url, nil) // want `should rewrite http.NewRequest to http.NewRequestWithContext or http.NewRequest and \(\*Request\).WithContext`
		req13, _ := http.NewRequest(http.MethodPost, url, nil) // want `should rewrite http.NewRequest to http.NewRequestWithContext or http.NewRequest and \(\*Request\).WithContext`
		return req12, req13
	}()
	cli.Do(req14)
	cli.Do(req15)
}
