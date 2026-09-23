package main

import (
	"fmt"

	"github.com/FilipeOliveira65/lookana/internal/monitor"
)

func main() {
	dnsList := []string{
		"google.com", "facebook.com", "spaces.live.com",
	}

	var counter int = 0

	for counter < len(dnsList) {
		var dnsToAnalyse = dnsList[counter]

		var dnsConsultResult = monitor.PingMonitor(dnsToAnalyse)

		fmt.Printf("Service name: %s\nService status: %s\n------------\n", dnsConsultResult.ServiceName, dnsConsultResult.ServiceStatus)

		counter += 1
	}
}
