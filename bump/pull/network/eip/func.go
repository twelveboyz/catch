package eip

import "Catch/bump/internal/mapping"

func Convert(s string) string {
	switch s {
	case "":
		return "未绑定"
	case "elb":
		return mapping.EIP_BindType_EIP
	case "snat":
		return mapping.EIP_BindType_SNAT

	case "ecs":
		return mapping.EIP_BindType_ECS
	default:
		return "NotFount=" + s
	}

}

func JoinPoolInfo(resources []*ElasticIP, pool string) {
	for _, resource := range resources {
		resource.Pool = pool
	}
}
