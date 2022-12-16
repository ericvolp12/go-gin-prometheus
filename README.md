# go-gin-prometheus
[![](https://godoc.org/github.com/ericvolp12/go-gin-prometheus?status.svg)](https://godoc.org/github.com/ericvolp12/go-gin-prometheus) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Gin Web Framework Prometheus metrics exporter

## Installation

`$ go get github.com/ericvolp12/go-gin-prometheus`

## Usage

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ericvolp12/go-gin-prometheus"
)

func main() {
	r := gin.New()

	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello world!")
	})

	r.Run(":29090")
}
```

See the [example.go file](https://github.com/ericvolp12/go-gin-prometheus/blob/main/example/example.go)
