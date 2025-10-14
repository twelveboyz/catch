package pfs

import (
	"Catch/bump/internal/mapping"
)

func Convert(s string) string {
	switch s {
	case "pfs":
		return mapping.PFS_PFS
	case "dpfs":
		return mapping.PFS_DPFS
	default:
		return "NotFound"
	}
}

func JoinPoolInfo(resources []*ParallelFileStorage, pool string) {
	for _, resource := range resources {
		resource.Pool = pool
	}
}
