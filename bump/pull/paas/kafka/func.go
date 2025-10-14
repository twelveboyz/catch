package kafka

import (
	"Catch/bump/internal/mapping"
)

func convert(s string) string {
	switch s {
	case "标准版":
		return mapping.KAFKA_STANDARD
	case "k8s":
		return mapping.KAFKA_K8S
	case "ekafka":
		return mapping.KAFKA_EKAFKA
	default:
		return "NotFound"
	}
}
