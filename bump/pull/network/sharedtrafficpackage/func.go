package sharedtrafficpackage

import (
	"Catch/bump/internal/mapping"
)

func Convert(s string) string {
	switch s {
	case "FULL_TIME":
		return mapping.STP_FULL_TIME
	default:
		return "NotFound"
	}
}
