package models

import (
	"time"

    "gorm.io/gorm" 
)

type ServiceData struct {
	gorm.Model 
	ServiceName string 
	ServiceDns string 
	ConsultFrenquencyInMin *uint
	LastTimeOffline *time.Time
	LastDnsCheck time.Time
}