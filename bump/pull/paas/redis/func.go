package redis

func convert(s string) string {
	switch s {
	case "SINGLE":
		return "单副本"
	case "MASTER_SLAVE":
		return "标准版"
	case "ECLUSTER":
		return "集群企业版"
	case "CLUSTER":
		return "集群社区版"
	case "BC_CLUSTER":
		return "集群代理版"
	default:
		return "NotFound"
	}
}
