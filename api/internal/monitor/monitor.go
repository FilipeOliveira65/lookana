package monitor

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/FilipeOliveira65/lookana/internal/models"
)

func PingMonitor(websiteDns string) models.PingResult {
	dnsToConsult := websiteDns

	var cmd *exec.Cmd

	cmd = exec.Command("ping", "-c", "3", dnsToConsult)

	pingResult, err := cmd.CombinedOutput()
	if err != nil {
		var pingOutput = string(pingResult)

		if !strings.Contains(pingOutput, "No address associated with hostname") {
			fmt.Println("Got an error here", err.Error())
		}
	}

	pingReturn := make(map[string]any)

	pingReturn["service-name"] = websiteDns

	if strings.Contains(string(pingResult), "No address associated with hostname") {
		return models.PingResult{
			ServiceName:   websiteDns,
			ServiceStatus: "offline",
		}
	}

	return models.PingResult{
		ServiceName:   websiteDns,
		ServiceStatus: "online",
	}
}
