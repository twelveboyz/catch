package wafProfessional

import (
	"fmt"
)

func BandwidthCount(i int64) string {
	i = (i * 50) + 100
	return fmt.Sprintf("%dMb", i)
}
