package eos

import (
	"Catch/bump/internal/mapping"
)

func PoolConvert(s string) string {
	switch s {
	case "guangzhou1":
		return mapping.EOS_POOL_GUANGZHOU1
	case "dongguan1":
		return mapping.EOS_POOL_DONGGUAN1
	case "dongguan7":
		return mapping.EOS_POOL_DONGGUAN7
	case "huanan1":
		return mapping.EOS_POOL_HUANAN1
	case "wuxi5":
		return mapping.EOS_POOL_WUXI5
	case "huhehaote1":
		return mapping.EOS_POOL_HUHEHAOTE1
	case "huhehaote6":
		return mapping.EOS_POOL_HUHEHAOTE6

	default:
		return "NotFount:" + s
	}
}

func JoinPoolInfo(resources []*ElasticObjectStorage, pool string) {
	for _, resource := range resources {
		resource.Pool = pool
	}
}
