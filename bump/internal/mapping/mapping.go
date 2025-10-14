package mapping

const (

	/*----------------------  资源池 ----------------------*/
	POOL_CIDC_RP_26 = "华南-广州3"
	POOL_CIDC_RP_25 = "华东-苏州"
	POOL_CIDC_RP_48 = "华北-呼和浩特"

	POOL_GUANGZHOU3 = "CIDC-RP-26"
	POOL_SUZHOU     = "CIDC-RP-25"
	POOL_HUHEHAOTE  = "CIDC-RP-48"

	/*----------------------  可用区 ----------------------*/
	AZ_DGJD           = "可用区1"
	AZ_N020_GD_GZNJ01 = "可用区2"
	AZ_N020_GD_GZFH01 = "可用区3"
	AZ_N020_GD_GZNJ04 = "可用区4"
	AZ_DongGuan       = "可用区1"
	AZ_NanJi          = "可用区2"
	AZ_FengHuang      = "可用区3"

	AZ_WXJD            = "可用区1"
	AZ_N0512_JS_SZFH01 = "可用区3"

	AZ_N0471_NMG_HHHT01 = "可用区1"

	/*----------------------  容器镜像服务  ----------------------*/
	CIS_Standard = "标准版"

	/*----------------------  系统盘  ----------------------*/
	SYSTEMDISK_SSD                        = "SSD"
	SYSTEMDISK_HIGHPERFORMANCE_YC         = "高性能型-云创版"
	SYSTEMDISK_HIGHPERFORMANCE            = "高性能型"
	SYSTEMDISK_PERFORMANCEOPTIMIZATION    = "性能优化型"
	SYSTEMDISK_PERFORMANCEOPTIMIZATION_YC = "性能优化型-云创版"

	/*----------------------  云硬盘  ----------------------*/
	EBS_SSD      = "高性能型"
	EBS_SSDEBS   = "性能优化型"
	EBS_SSDYC    = "高性能型-云创版"
	EBS_SSDEBSYC = "性能优化性-云创版"
	/*----------------------  对象存储  ----------------------*/
	//EOS 华南-广州2
	EOS_EOS_POOL_ID_guangzhou1 = "CIDC-RP-02"
	EOS_POOL_ID_guangzhou1     = "CIDC-RP-25"

	//EOS 华北-呼和浩特
	EOS_EOS_POOL_ID_huhehaote1 = "CIDC-RP-48"
	EOS_POOL_ID_huhehaote1     = "CIDC-RP-48"

	//EOS 华北-呼和浩特2
	EOS_EOS_POOL_ID_huhehaote6 = "CIDC-RP-632"
	EOS_POOL_ID_huhehaote6     = "CIDC-RP-48"

	EOS_POOL_GUANGZHOU1 = "华南-广州2"
	EOS_POOL_DONGGUAN1  = "华南-广州3"
	EOS_POOL_DONGGUAN7  = "华南-广州5"
	EOS_POOL_HUANAN1    = "华中-长沙1"
	EOS_POOL_WUXI5      = "华东-无锡"
	EOS_POOL_HUHEHAOTE1 = "华北-呼和浩特"
	EOS_POOL_HUHEHAOTE6 = "华北-呼和浩特2"
	/*----------------------  文件存储  ----------------------------*/
	EFS_NAS  = "容量型"
	EFS_PNAS = "性能型"
	/*----------------------  并行文件存储  ----------------------------*/
	PFS_PFS  = "标准型"
	PFS_DPFS = "极速型"

	/*----------------------  负载均衡  ----------------------------*/
	SLB_FLAVOR_2  = "优享型I"
	SLB_FLAVOR_3  = "优享型II"
	SLB_FLAVOR_4  = "高端型I"
	SLB_FLAVOR_5  = "高端型II"
	SLB_FLAVOR_6  = "旗舰型"
	SLB_FLAVOR_21 = "性能保障型"
	SLB_FLAVOR_30 = "LCU"

	/*----------------------  负载均衡  ----------------------------*/
	EIP_BindType_EIP  = "弹性负载均衡"
	EIP_BindType_SNAT = "NAT网关"
	EIP_BindType_ECS  = "云主机"
	/*----------------------  NAT网关  ----------------------*/
	NATGATEWAY_SMALL       = "小型"
	NATGATEWAY_MIDDLE      = "中型"
	NATGATEWAY_LARGE       = "大型"
	NATGATEWAY_ULTRA_LARGE = "超大型"
	/*----------------------  共享流量包  ----------------------*/
	STP_FULL_TIME = "全时"
	/*----------------------  redis  ----------------------*/
	REDIS_SINGLE       = "单副本"
	REDIS_MASTER_SLAVE = "标准版"
	REDIS_ECLUSTER     = "集群企业版"
	REDIS_CLUSTER      = "集群社区版"
	REDIS_BC_CLUSTER   = "集群代理版"
	/*----------------------  rabbit  ---------------------*/
	RABBIT_region1 = "可用区一"
	RABBIT_region2 = "可用区二"
	RABBIT_region3 = "可用区三"
	/*----------------------  kafka  ---------------------*/
	KAFKA_STANDARD = "标准版"
	KAFKA_K8S      = "专业版"
	KAFKA_EKAFKA   = "旗舰版"
	/*----------------------  HBase  ---------------------*/
	HBASE_STORAGECONFIG_TYPE_0 = "HDD"
	HBASE_STORAGECONFIG_TYPE_1 = "SSD"
)
