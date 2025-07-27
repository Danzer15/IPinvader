// Ini Untuk Resolver
package domain

import (
	"net"
)

func GetDomain(ip string) string {
	name, err := net.LookupAddr(ip)
	if err != nil || len(name) == 0 {
		return "Tidak ditemukan"
	}
	return name[0]
}
