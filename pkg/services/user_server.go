package services

import (
	"net"
	"math/rand"
)

// Генерация n случайных IPv4 адресов
func generateIPs(n int) []net.IP {
	ips := make([]net.IP, n)
	for i := 0; i < n; i++ {
		ip := make(net.IP, 4)
		for j := range ip {
			ip[j] = byte(rand.Intn(256))
		}
		ips[i] = ip
	}
	return ips
}
