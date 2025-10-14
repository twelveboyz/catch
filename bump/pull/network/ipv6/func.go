package ipv6

func JoinPoolInfo(resources []*IPv6, pool string) {
	for _, resource := range resources {
		resource.Pool = pool
	}
}
