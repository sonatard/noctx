package http_client

import (
	"net/http"
)

func _() {
	const url = "http://example.com"
	cli := &http.Client{}

	http.Get(url) // want `net/http\.Get must not be called. use net/http\.NewRequestWithContext and \(\*net/http.Client\)\.Do\(\*http.Request\)`
	_ = http.Get  // OK
	f := http.Get // OK
	f(url)        // want `net/http\.Get must not be called. use net/http\.NewRequestWithContext and \(\*net/http.Client\)\.Do\(\*http.Request\)`

	http.Head(url)          // want `net/http\.Head must not be called. use net/http\.NewRequestWithContext and \(\*net/http.Client\)\.Do\(\*http.Request\)`
	http.Post(url, "", nil) // want `net/http\.Post must not be called. use net/http\.NewRequestWithContext and \(\*net/http.Client\)\.Do\(\*http.Request\)`
	http.PostForm(url, nil) // want `net/http\.PostForm must not be called. use net/http\.NewRequestWithContext and \(\*net/http.Client\)\.Do\(\*http.Request\)`

	cli.Get(url) // want `\(\*net/http\.Client\)\.Get must not be called. use \(\*net/http.Client\)\.Do\(\*http.Request\)`
	_ = cli.Get  // OK
	m := cli.Get // OK
	m(url)       // want `\(\*net/http\.Client\)\.Get must not be called. use \(\*net/http.Client\)\.Do\(\*http.Request\)`

	cli.Head(url)          // want `\(\*net/http\.Client\)\.Head must not be called. use \(\*net/http.Client\)\.Do\(\*http.Request\)`
	cli.Post(url, "", nil) // want `\(\*net/http\.Client\)\.Post must not be called. use \(\*net/http.Client\)\.Do\(\*http.Request\)`
	cli.PostForm(url, nil) // want `\(\*net/http\.Client\)\.PostForm must not be called. use \(\*net/http.Client\)\.Do\(\*http.Request\)`
}
