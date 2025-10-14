package cis

import (
	"Catch/bump/internal/mapping"
)

func Convert(s string) string {
	switch s {
	case "Standard":
		return mapping.CIS_Standard
	default:
		return "NotFound"
	}
}
