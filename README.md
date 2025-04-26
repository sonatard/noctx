# noctx

![](https://github.com/sonatard/noctx/workflows/CI/badge.svg)

`noctx` finds function calls without context.Context.

If you are using net/http package and sql/database package, you should use noctx.
Passing `context.Context` enables library user to cancel request, getting trace information and so on.

## Usage

### noctx with go vet

go vet is a Go standard tool for analyzing source code.

1. Install noctx.
```sh
$ go install github.com/sonatard/noctx/cmd/noctx@latest
```

2. Execute noctx
```sh
$ go vet -vettool=`which noctx` main.go
./main.go:6:11: net/http.Get must not be called
```

### noctx with golangci-lint

golangci-lint is a fast Go linters runner.

1. Install golangci-lint.
[golangci-lint - Install](https://golangci-lint.run/usage/install/)

2. Setup .golangci.yml
```yaml:
# Add noctx to enable linters.
linters:
  enable:
    - noctx

# Or enable-all is true.
linters:
  enable-all: true
  disable:
   - xxx # Add unused linter to disable linters.
```

3. Execute noctx
```sh
# Use .golangci.yml
$ golangci-lint run

# Only execute noctx
golangci-lint run --enable-only noctx
```

## net/http package
### Detection rules

- Executing the following functions
https://github.com/sonatard/noctx/blob/9a514098df3f8a88e0fd6949320c4e0aa51b520c/ngfunc/main.go#L15-L26

### How to fix

- use `http.NewRequestWithContext` function instead of using `http.NewRequest` function.
- Send http request using `(*http.Client).Do(*http.Request)` method.

If your library already provides functions that don't accept context, you define a new function that accepts context and make the existing function a wrapper for a new function.

```go
// Before fix code
// Sending an HTTP request but not accepting context
func Send(body io.Reader)  error {
    req,err := http.NewRequest(http.MethodPost, "http://example.com", body)
    if err != nil {
        return err
    }
    _, err = http.DefaultClient.Do(req)
    if err != nil{
        return err
    }

    return nil
}
```

```go
// After fix code
func Send(body io.Reader) error {
    // Pass context.Background() to SendWithContext
    return SendWithContext(context.Background(), body)
}

// Sending an HTTP request and accepting context
func SendWithContext(ctx context.Context, body io.Reader) error {
    // Change NewRequest to NewRequestWithContext and pass context to it
    req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://example.com", body)
    if err != nil {
        return err
    }
    _, err = http.DefaultClient.Do(req)
    if err != nil {
        return err
    }

    return nil
}
```

### Detection sample

https://github.com/sonatard/noctx/blob/9a514098df3f8a88e0fd6949320c4e0aa51b520c/testdata/src/http_client/http_client.go#L11
https://github.com/sonatard/noctx/blob/9a514098df3f8a88e0fd6949320c4e0aa51b520c/testdata/src/http_request/http_request.go#L17

### Reference

- [net/http - NewRequest](https://pkg.go.dev/net/http#NewRequest)
- [net/http - NewRequestWithContext](https://pkg.go.dev/net/http#NewRequestWithContext)
- [net/http - Request.WithContext](https://pkg.go.dev/net/http#Request.WithContext)
- 
## database/sqlpackage
### Detection rules

- Executing the following functions

### Detection sample


### Reference
- [database/sql](https://pkg.go.dev/database/sql)
