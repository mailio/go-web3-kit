# Go Web3 Kit

Work in progress...

## Prerequisites

install Swag:
- `go install github.com/swaggo/swag/cmd/swag@latest`

init swag everytime new http API is added:
```go
swag init --parseDependency
```

## Generating API Documentation

Check swagger docs by directing a browser to: `/swagger/index.html`

## Projects

- [gin-swagger](https://github.com/swaggo/gin-swagger)
- [gin-gonic](https://github.com/gin-gonic/gin)
- [go-resty](https://github.com/go-resty/resty)
- [go-kit/log](https://github.com/go-kit/log)