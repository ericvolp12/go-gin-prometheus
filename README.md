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


$ # In another window (note the 10,000 VUs here is a LOT so you can knock that down to like 500)
$ k6 run --vus 10000 --duration 30s k6_test.js

  execution: local
     script: k6_test.js
     output: -

  scenarios: (100.00%) 1 scenario, 10000 max VUs, 1m0s max duration (incl. graceful stop):
           * default: 10000 looping VUs for 30s (gracefulStop: 30s)


running (0m35.0s), 00000/10000 VUs, 126366 complete and 0 interrupted iterations
default ✓ [======================================] 10000 VUs  30s

     data_received..................: 17 MB  495 kB/s
     data_sent......................: 10 MB  293 kB/s
     http_req_blocked...............: avg=7.27ms  min=520ns   med=940ns   max=187.2ms  p(90)=3.06µs p(95)=86.21ms
     http_req_connecting............: avg=6.7ms   min=0s      med=0s      max=187.17ms p(90)=0s     p(95)=81.35ms
     http_req_duration..............: avg=2.49s   min=53.24µs med=2.48s   max=5.06s    p(90)=4.49s  p(95)=4.74s
       { expected_response:true }...: avg=2.49s   min=53.24µs med=2.48s   max=5.06s    p(90)=4.49s  p(95)=4.74s
     http_req_failed................: 0.00%  ✓ 0          ✗ 126366
     http_req_receiving.............: avg=17.66µs min=5.94µs  med=16.35µs max=60.61ms  p(90)=21.7µs p(95)=23.65µs
     http_req_sending...............: avg=1.37ms  min=2.7µs   med=4.98µs  max=80.33ms  p(90)=9.22µs p(95)=5.34ms
     http_req_tls_handshaking.......: avg=0s      min=0s      med=0s      max=0s       p(90)=0s     p(95)=0s
     http_req_waiting...............: avg=2.49s   min=39.39µs med=2.48s   max=5.01s    p(90)=4.49s  p(95)=4.74s
     http_reqs......................: 126366 3611.73162/s
     iteration_duration.............: avg=2.49s   min=73.9µs  med=2.49s   max=5.13s    p(90)=4.5s   p(95)=4.75s
     iterations.....................: 126366 3611.73162/s
     vus............................: 34     min=34       max=10000
     vus_max........................: 10000  min=10000    max=10000
```

An example output from the `/metrics` endpoint after executing a test looks like:

```
# HELP gin_request_duration_seconds The HTTP request latencies in seconds.
# TYPE gin_request_duration_seconds histogram
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="0.0001"} 25
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="0.005"} 134
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="0.01"} 249
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="0.025"} 653
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="0.05"} 1280
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="0.1"} 2556
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="0.25"} 6413
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="0.5"} 12748
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="1"} 25314
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="2.5"} 63511
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="5"} 126365
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="10"} 126366
gin_request_duration_seconds_bucket{code="200",method="GET",url="/",le="+Inf"} 126366
gin_request_duration_seconds_sum{code="200",method="GET",url="/"} 314792.871235347
gin_request_duration_seconds_count{code="200",method="GET",url="/"} 126366
# HELP gin_request_size_bytes The HTTP request sizes in bytes.
# TYPE gin_request_size_bytes summary
gin_request_size_bytes_sum 7.961058e+06
gin_request_size_bytes_count 126366
# HELP gin_requests_total How many HTTP requests processed, partitioned by status code and HTTP method.
# TYPE gin_requests_total counter
gin_requests_total{code="200",handler="main.main.func1",host="localhost:29090",method="GET",url="/"} 126366
# HELP gin_response_size_bytes The HTTP response sizes in bytes.
# TYPE gin_response_size_bytes summary
gin_response_size_bytes_sum 1.769124e+06
gin_response_size_bytes_count 126366
```
