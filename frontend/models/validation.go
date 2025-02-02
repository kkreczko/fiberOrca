package models

import (
	"fmt"
	"regexp"
)

func ValidateIP(IPString string) error {
	match, _ := regexp.MatchString("^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$", IPString)
	if !match && IPString != "" {
		return fmt.Errorf("invalid IP address")
	}
	return nil
}

func ValidatePort(portString string) error {
	match, _ := regexp.MatchString("^[0-9]{1,4}$", portString)
	if !match && portString != "" {
		return fmt.Errorf("invalid port")
	}
	return nil
}

func ValidateTime(timeString string) error {
	match, _ := regexp.MatchString("^(\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}) $", timeString)
	if !match && timeString != "" {
		return fmt.Errorf("invalid time")
	}
	return nil
}
