package hbase

import (
	"Catch/bump/internal/mapping"
)

func convert(s string) string {
	switch s {
	case "0":
		return mapping.HBASE_STORAGECONFIG_TYPE_0
	case "1":
		return mapping.HBASE_STORAGECONFIG_TYPE_1
	default:
		return "NotFound"
	}
}
