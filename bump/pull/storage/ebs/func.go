package ebs

import (
	"Catch/bump/internal/mapping"
)

func Convert(s string) string {
	switch s {
	case "ssd":
		return mapping.EBS_SSD
	case "ssdebs":
		return mapping.EBS_SSDEBS
	case "ssdyc":
		return mapping.EBS_SSDYC
	case "ssdebsyc":
		return mapping.EBS_SSDEBSYC
	default:
		return "NotFound:" + s
	}
}

func isShareConvert(s string) string {
	switch s {
	case "true":
		return "共享"
	case "false":
		return "非共享"
	default:
		return "NotFount:" + s
	}
}

func JoinPoolInfo(resources []*ElasticBlockStorage, pool string) {
	for _, resource := range resources {
		resource.Pool = pool
	}
}
