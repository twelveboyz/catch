package natgateway

import (
	"Catch/bump/internal/mapping"
)

func Convert(s string) string {
	switch s {
	case "small":
		return mapping.NATGATEWAY_SMALL
	case "middle":
		return mapping.NATGATEWAY_MIDDLE
	case "large":
		return mapping.NATGATEWAY_LARGE
	case "ultra-large":
		return mapping.NATGATEWAY_ULTRA_LARGE
	default:
		return "NotFound:" + s
	}
}

func JoinPoolInfo(resources []*NATGateway, pool string) {
	for _, resource := range resources {
		resource.Pool = pool
	}
}
