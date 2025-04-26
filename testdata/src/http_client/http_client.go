package http_client

import (
	"net/http"
)

func _() {
	const url = "http://example.com"
	cli := &http.Client{}

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
}
