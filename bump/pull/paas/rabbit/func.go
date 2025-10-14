package rabbit

import (
	"Catch/bump/internal/mapping"
)

func convert(s string) string {
	switch s {
	case "region1":
		return mapping.RABBIT_region1
	case "region2":
		return mapping.RABBIT_region2
	case "region3":
		return mapping.RABBIT_region3
	default:
		return "NotFound"
	}
}
