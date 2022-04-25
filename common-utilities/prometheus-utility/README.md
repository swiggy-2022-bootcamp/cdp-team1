# Introduction
- A prometheus utility project that can be used by any other microservice by just importing it as a go module.
- Prometheus is an open-source systems monitoring and alerting toolkit originally built at SoundCloud.
- Prometheus collects and stores its metrics as time series data, i.e. metrics information is stored with the timestamp at which it was recorded, alongside optional key-value pairs called labels.
- Metrics defined in the project currently are:
  - response_status
  - http_requests_total
  - http_response_time_seconds
  
# Usage guide

- prometheus-utility library can be directly imported with the following command

  `go get -u github.com/swiggy-2022-bootcamp/cdp-team1/common-utilities/prometheus-utility`
- Import the library

  `import (
  prometheusUtility "github.com/swiggy-2022-bootcamp/cdp-team1/common-utilities/prometheus-utility"
  )`
- Create a route exposing the /metrics endpoint.

  `router.GET("/metrics", prometheusUtility.PrometheusHandler())`
- Register the metrics in your main.go file

  `func init() {
  prometheusUtility.RegisterMetrics()
  }`
- Use the middleware to intercept requests and collect metrics

  `server.Use(prometheusUtility.PrometheusMiddleware())`