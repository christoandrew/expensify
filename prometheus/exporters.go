package prometheus

import "github.com/prometheus/client_golang/prometheus"

var ExpensesRequestCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "expenses_request_count",
	Help: "Number of requests expenses endpoint",
})
