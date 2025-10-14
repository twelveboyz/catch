package loadbalance

func JoinPoolInfo(resources []*LoadBalance, pool string) {
	for _, resource := range resources {
		resource.Pool = pool
	}
}
