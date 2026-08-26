package models

type PingResult struct {
	ServiceName   string `json:"service-name"`
	ServiceStatus string `json:"service-status"`
}
