package main

import (
	"math/rand"
	"time"

	ginprometheus "github.com/ericvolp12/go-gin-prometheus"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	/*	// Optional custom metrics list
		customMetrics := []*ginprometheus.Metric{
			&ginprometheus.Metric{
				ID:	"1234",				// optional string
				Name:	"test_metric",			// required string
				Description:	"Counter test metric",	// required string
				Type:	"counter",			// required string
			},
			&ginprometheus.Metric{
				ID:	"1235",				// Identifier
				Name:	"test_metric_2",		// Metric Name
				Description:	"Summary test metric",	// Help Description
				Type:	"summary", // type associated with prometheus collector
			},
			// Type Options:
			//	counter, counter_vec, gauge, gauge_vec,
			//	histogram, histogram_vec, summary, summary_vec
		}
		p := ginprometheus.NewPrometheus("gin", customMetrics)
	*/

	p := ginprometheus.NewPrometheus("gin", &ginprometheus.DefaultMetricOverrides{
		// Create custom buckets that extends the low-range down to 100 microseconds
		RequestDurationSecondsBuckets: &[]float64{
			0.0001, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10,
		},
	})

	rand.Seed(time.Now().UnixMilli())

	p.Use(r)
	r.GET("/", func(c *gin.Context) {
		// Generate a random latency from 0-5000 milliseconds
		pauseForMS := rand.Float64() * 5000
		time.Sleep(time.Millisecond * time.Duration(pauseForMS))
		c.JSON(200, "Hello world!")
	})

	r.Run(":29090")
}
