package elk

import (
	"fmt"
	"regexp"
	"strconv"
)

type LogData struct {
	IP        string `json:"ip"`
	Timestamp string `json:"timestamp"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	Status    int    `json:"status"`
	Bytes     int    `json:"bytes"`
}

// Parse a single log line into a structured format
func parseLogLine(logLine string) (*LogData, error) {
	regex := regexp.MustCompile(`(?P<ip>\S+) \S+ \S+ \[(?P<timestamp>[^\]]+)\] "(?P<method>\S+) (?P<path>\S+) \S+" (?P<status>\d+) (?P<bytes>\d+)`)
	match := regex.FindStringSubmatch(logLine)

	if match == nil {
		return nil, fmt.Errorf("log line does not match expected format")
	}

	// Extract values using named groups
	groupNames := regex.SubexpNames()
	logData := &LogData{}

	for i, name := range groupNames {
		if i == 0 || name == "" {
			continue
		}

		value := match[i]
		switch name {
		case "ip":
			logData.IP = value
		case "timestamp":
			logData.Timestamp = value
		case "method":
			logData.Method = value
		case "path":
			logData.Path = value
		case "status":
			logData.Status = atoi(value)
		case "bytes":
			logData.Bytes = atoi(value)
		}
	}

	return logData, nil
}

// Helper function to convert string to int
func atoi(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}
