package ironic

func JoinPoolInfo(ironics []*Ironic, pool string) {
	for _, ironic := range ironics {
		ironic.Pool = pool
	}
}
