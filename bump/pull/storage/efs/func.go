package efs

import (
	"Catch/bump/internal/mapping"
	"errors"
	"fmt"
	"strings"
)

func Convert(s string) string {
	switch s {
	case "nas":
		return mapping.EFS_NAS
	case "pnas":
		return mapping.EFS_PNAS
	default:
		return "NotFound"
	}
}

func JoinPoolInfo(resources []*ElasticFileStorage, pool string) {
	for _, resource := range resources {
		resource.Pool = pool
	}
}

func parseIPAndSplit(s string, split string) (string, string, error) {
	slice := strings.Split(s, split)

	if len(slice) != 2 {
		return "", "", errors.New("length ! = 2")
	}

	return slice[0], slice[1], nil

}

func MountInfoFormat(exportLocation string) (string, error) {
	ipv4, ipv6, err := parseIPAndSplit(exportLocation, "|")
	if err != nil {
		return "", err
	}

	var ipv4Sum string
	ipv4s, ipv4Suffix, err := parseIPAndSplit(ipv4, ":/")
	if err != nil {
		return "", err
	}

	ipv4Slice := strings.Split(ipv4s, ",")
	for _, ip := range ipv4Slice {
		ipv4Sum = ipv4Sum + fmt.Sprintf("%s:/%s\n", ip, ipv4Suffix)
	}

	var ipv6Sum string
	ipv6s, ipv6Suffix, err := parseIPAndSplit(ipv6, ":/")
	if err != nil {
		return "", err
	}

	ipv6Slice := strings.Split(ipv6s, ",")
	for _, ip := range ipv6Slice {
		ipv6Sum = ipv6Sum + fmt.Sprintf("%s:/%s\n", ip, ipv6Suffix)
	}

	return strings.TrimRight(ipv4Sum+ipv6Sum, "\n"), nil
}
