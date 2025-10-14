package mapping

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func AZConvert(s string) string {
	switch s {
	case "DGJD":
		return AZ_DGJD
	case "N020-GD-GZNJ01":
		return AZ_N020_GD_GZNJ01
	case "N020-GD-GZFH01":
		return AZ_N020_GD_GZFH01
	case "N020-GD-GZNJ04":
		return AZ_N020_GD_GZNJ04
	case "dongguan":
		return AZ_DongGuan
	case "nanji":
		return AZ_NanJi
	case "fenghuang":
		return AZ_FengHuang
	case "N0471-NMG-HHHT01":
		return AZ_N0471_NMG_HHHT01
	case "N0512-JS-SZFH01":
		return AZ_N0512_JS_SZFH01
	case "WXJD":
		return AZ_WXJD
	case "可用区1":
		return "可用区1"
	case "可用区2":
		return "可用区2"
	case "可用区3":
		return "可用区3"
	default:
		return fmt.Sprintf("NotFound:%s", s)
	}
}

func DiskConvert(s string) string {
	switch s {
	case "local":
		return SYSTEMDISK_SSD
	case "highPerformance":
		return SYSTEMDISK_HIGHPERFORMANCE
	case "highPerformanceyc":
		return SYSTEMDISK_HIGHPERFORMANCE_YC
	case "performanceOptimization":
		return SYSTEMDISK_PERFORMANCEOPTIMIZATION
	case "performanceOptimizationyc":
		return SYSTEMDISK_PERFORMANCEOPTIMIZATION_YC
	default:
		return fmt.Sprintf("NotFound:%s", s)
	}
}

func PoolNameToCodeConvert(s string) string {
	switch s {
	case POOL_CIDC_RP_26:
		return POOL_GUANGZHOU3
	case POOL_CIDC_RP_25:
		return POOL_SUZHOU
	case POOL_CIDC_RP_48:
		return POOL_HUHEHAOTE
	default:
		logrus.Warnf("NotFound Pool:%s\n", s)
		return fmt.Sprintf("NotFound:%s", s)
	}
}

func PoolCodeToNameConvert(s string) string {
	switch s {
	case POOL_GUANGZHOU3:
		return POOL_CIDC_RP_26
	case POOL_SUZHOU:
		return POOL_CIDC_RP_25
	case POOL_HUHEHAOTE:
		return POOL_CIDC_RP_48
	default:
		logrus.Warnf("NotFound Pool:%s\n", s)
		return fmt.Sprintf("NotFound:%s", s)

	}
}
