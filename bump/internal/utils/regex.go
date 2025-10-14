package utils

import (
	"net"
	"regexp"
	"strconv"
	"strings"
)

func CidrToMask(s string) string {
	r := regexp.MustCompile(`/\d{1,3}$`)
	rCidr := r.FindString(s)
	rCidr = strings.ReplaceAll(rCidr, "/", "")

	v4cidrInt, _ := strconv.Atoi(rCidr)
	mask := net.IP(net.CIDRMask(v4cidrInt, 32)).String()

	return mask

}
