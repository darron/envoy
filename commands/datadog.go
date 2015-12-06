package commands

import (
	"fmt"
	"github.com/PagerDuty/godspeed"
	"github.com/zorkian/go-datadog-api"
	"time"
)

// SendStats sends statistics to the Datadog API using either Dogstatsd or the
// REST API.
func SendStats(nodes []Node, start time.Time) {
	elapsed := time.Since(start)
	hostCount := int(len(nodes))
	milliseconds := int64(elapsed / time.Millisecond)
	var tags []string
	envTag := fmt.Sprintf("environment:%s", ChefEnvironment)
	tags = append(tags, envTag)
	if DogStatsd {
		statsd, _ := godspeed.New(DogStatsdAddress, godspeed.DefaultPort, false)
		defer statsd.Conn.Close()
		statsd.Gauge("envoy.time", float64(milliseconds), tags)
		statsd.Gauge("envoy.hosts", float64(hostCount), tags)
		Log(fmt.Sprintf("create: dogstatsd='true' time='%s' hosts='%d'", elapsed, hostCount), "info")
	}
	if DatadogAPIKey != "" && DatadogAPPKey != "" {
		client := datadog.NewClient(DatadogAPIKey, DatadogAPPKey)
		timestamp := float64(time.Now().Unix())
		err := client.PostMetrics([]datadog.Metric{
			{
				Metric: "envoy.time",
				Points: []datadog.DataPoint{
					{timestamp, float64(milliseconds)},
				},
				Tags: tags,
			},
			{
				Metric: "envoy.hosts",
				Points: []datadog.DataPoint{
					{timestamp, float64(hostCount)},
				},
				Tags: tags,
			},
		})
		if err != nil {
			Log("Something went wrong.", "info")
		}
	}
}
