package handlers


import (
"fmt"
"net"
"strings"
)

// CommandHandler обрабатывает команды игры
type CommandHandler struct {
	IPs []net.IP
}

func (h *CommandHandler) Scan(mask string) []net.IP {
	var matched []net.IP
	for _, ip := range h.IPs {
		if strings.HasPrefix(ip.String(), mask) {
			matched = append(matched, ip)
		}
	}
	return matched
}

func (h *CommandHandler) Ping(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false
	}

	for _, existingIP := range h.IPs {
		if existingIP.Equal(ip) {
			return true
		}
	}
	return false
}

func (h *CommandHandler) Trace(ipStr string) []string {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return nil
	}

	var neighbors []string
	ipBytes := ip.To4()

	for _, existingIP := range h.IPs {
		existingBytes := existingIP.To4()
		if existingBytes[0] == ipBytes[0] && existingBytes[1] == ipBytes[1] &&
			!(existingBytes[2] == ipBytes[2] && existingBytes[3] == ipBytes[3]) {
			neighbor := fmt.Sprintf("%d.%d.*.*", existingBytes[0], existingBytes[1])
			neighbors = append(neighbors, neighbor)
			if len(neighbors) >= 2 {
				break
			}
		}
	}

	return neighbors
}

func (h *CommandHandler) RemoveIP(ip net.IP) {
	for i, existingIP := range h.IPs {
		if existingIP.Equal(ip) {
			h.IPs = append(h.IPs[:i], h.IPs[i+1:]...)
			break
		}
	}
}