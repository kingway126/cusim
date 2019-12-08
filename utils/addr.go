package utils

import "strings"

func ParseAddr(host string) string {
	path := strings.Split(host, ":")
	return path[0]
}
