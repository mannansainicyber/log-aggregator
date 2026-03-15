package analyzer

import (
	"fmt"
	"log-ag/db"
	"strings"
)

type Alert struct {
	Type    string
	Service string
	Message string
	Time    string
}

func checkKeywords(log db.Log) *Alert {
	for _, keyword := range suspiciousKeywords {
		if strings.Contains(log.Message, keyword) {
			return &Alert{
				Type:    "SUSPICIOUS",
				Service: log.Service,
				Message: "contains keyword: " + keyword,
				Time:    log.Timestamp,
			}
		}
	}
	return nil
}

func checkRepeat(logs []db.Log) []Alert {
	counts := map[string]int{}

	for _, log := range logs {
		counts[log.Message]++
	}

	var alerts []Alert

	for message, count := range counts {
		if count >= repeatThreshold {
			alerts = append(alerts, Alert{
				Type:    "REPEAT",
				Message: fmt.Sprintf("'%s' appeared %d times", message, count),
			})
		}
	}
	return alerts
}

func checkBurst(logs []db.Log) []Alert {
	errorCount := 0

	for _, log := range logs {
		if log.Level == "error" {
			errorCount++
		}
	}

	var alerts []Alert
	if errorCount >= burstThreshold {
		alerts = append(alerts, Alert{
			Type:    "BURST",
			Message: fmt.Sprintf("%d errors detected in logs", errorCount),
		})
	}
	return alerts
}

func Analyze(logs []db.Log) []Alert {
	var alerts []Alert

	for _, log := range logs {
		if a := checkKeywords(log); a != nil {
			alerts = append(alerts, *a)
		}
	}

	b := checkRepeat(logs)
	c := checkBurst(logs)
	alerts = append(alerts, b...)
	alerts = append(alerts, c...)
	// append their results to alerts

	return alerts
}
