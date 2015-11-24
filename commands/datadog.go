package commands

import (
	"fmt"
	"github.com/PagerDuty/godspeed"
	"time"
)

func SendStats(nodes []Node, start time.Time) {
	elapsed := time.Since(start)
	if DogStatsd {
		hostCount := int(len(nodes))
		milliseconds := int64(elapsed / time.Millisecond)
		statsd, _ := godspeed.New(DogStatsdAddress, godspeed.DefaultPort, false)
		defer statsd.Conn.Close()
		tags := make([]string, 0)
		envTag := fmt.Sprintf("environment:%s", ChefEnvironment)
		tags = append(tags, envTag)
		statsd.Gauge("envoy.time", float64(milliseconds), tags)
		statsd.Gauge("envoy.hosts", float64(hostCount), tags)
		Log(fmt.Sprintf("create: dogstatsd='true' time='%s' hosts='%d'", elapsed, hostCount), "info")
	}
}
