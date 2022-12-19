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

## Testing

This package uses [`k6`](https://k6.io/) to generate a load against the example in order to test the output metrics.

An exmaple test should produce a random distribution of request times between 0 and 5 seconds with a median and mean response time around 2.5 seconds:

```shell
$ go run example/example.go

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /metrics                  --> github.com/ericvolp12/go-gin-prometheus.prometheusHandler.func1 (2 handlers)
[GIN-debug] GET    /                         --> main.main.func1 (2 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :29090


$ # In another window
$ k6 run --vus 50 --duration 30s k6_test.js

  scenarios: (100.00%) 1 scenario, 50 max VUs, 1m0s max duration (incl. graceful stop):
           * default: 50 looping VUs for 30s (gracefulStop: 30s)


running (0m33.9s), 00/50 VUs, 627 complete and 0 interrupted iterations
default ✓ [======================================] 50 VUs  30s

     data_received..................: 86 kB 2.5 kB/s
     data_sent......................: 51 kB 1.5 kB/s
     http_req_blocked...............: avg=25.71µs min=720ns   med=1.71µs  max=542.65µs p(90)=3.65µs  p(95)=247.82µs
     http_req_connecting............: avg=13.19µs min=0s      med=0s      max=432.97µs p(90)=0s      p(95)=137.28µs
     http_req_duration..............: avg=2.51s   min=16.93ms med=2.48s   max=4.98s    p(90)=4.57s   p(95)=4.72s
       { expected_response:true }...: avg=2.51s   min=16.93ms med=2.48s   max=4.98s    p(90)=4.57s   p(95)=4.72s
     http_req_failed................: 0.00% ✓ 0         ✗ 627
     http_req_receiving.............: avg=32.98µs min=15.37µs med=32.63µs max=69.37µs  p(90)=46.5µs  p(95)=50.42µs
     http_req_sending...............: avg=9.87µs  min=4µs     med=8.25µs  max=60.17µs  p(90)=13.47µs p(95)=25.48µs
     http_req_tls_handshaking.......: avg=0s      min=0s      med=0s      max=0s       p(90)=0s      p(95)=0s
     http_req_waiting...............: avg=2.51s   min=16.84ms med=2.48s   max=4.98s    p(90)=4.57s   p(95)=4.72s
     http_reqs......................: 627   18.503394/s
     iteration_duration.............: avg=2.51s   min=17.52ms med=2.48s   max=4.98s    p(90)=4.57s   p(95)=4.72s
     iterations.....................: 627   18.503394/s
     vus............................: 6     min=6       max=50
     vus_max........................: 50    min=50      max=50
```

An example output from the `/metrics` endpoint after executing a test looks like:

```
# HELP gin_request_duration_seconds The HTTP request latencies in seconds.
# TYPE gin_request_duration_seconds histogram
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="0.0001"} 0
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="0.005"} 0
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="0.01"} 0
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="0.025"} 2
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="0.05"} 6
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="0.1"} 9
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="0.25"} 24
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="0.5"} 60
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="1"} 116
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="2.5"} 318
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="5"} 627
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="10"} 627
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="+Inf"} 627
gin_request_duration_seconds_sum{code="200",method="GET",url="/"} 1573.7542368
gin_request_duration_seconds_count{code="200",method="GET",url="/"} 627
# HELP gin_request_size_bytes The HTTP request sizes in bytes.
# TYPE gin_request_size_bytes summary
gin_request_size_bytes_sum 39501
gin_request_size_bytes_count 627
# HELP gin_requests_total How many HTTP requests processed, partitioned by status code and HTTP method.
# TYPE gin_requests_total counter
gin_requests_total{code="200",handler="main.main.func1",host="localhost:29090",method="GET",url="/"} 627
# HELP gin_response_size_bytes The HTTP response sizes in bytes.
# TYPE gin_response_size_bytes summary
gin_response_size_bytes_sum 8778
gin_response_size_bytes_count 627
```
