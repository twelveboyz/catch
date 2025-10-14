package ecs

func JoinPoolInfo(ecss []*ECloudServer, pool string) {
	for _, ecs := range ecss {
		ecs.Pool = pool
	}
}
