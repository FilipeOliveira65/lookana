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

		// result := map[any]any{
		// 	"service-name":   fmt.Sprintf("%s", dnsToAnalyse),
		// 	"service-status": fmt.Sprintf("%s", dnsConsultResult),
		// }

		// for key, val := range result {
		// 	fmt.Printf("Service: %s | Status: %s\n", key, val)
		// }

		counter += 1
	}

}
