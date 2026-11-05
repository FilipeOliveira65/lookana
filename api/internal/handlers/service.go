package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FilipeOliveira65/lookana/internal/monitor"
)

type NewService struct {
	ServiceName string `json:"service-name"`
	ServiceDns  string `json:"service-dns"`
}

func CreateService(w http.ResponseWriter, r *http.Request) {
	var newService NewService

	err := json.NewDecoder(r.Body).Decode(&newService)

	if err != nil {
		http.Error(w, "Invalid JSON.", http.StatusBadRequest)
	}

	fmt.Println("Reached here")
	fmt.Println(newService)
}

func PingService(w http.ResponseWriter, r *http.Request) {
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
