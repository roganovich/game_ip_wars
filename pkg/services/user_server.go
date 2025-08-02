package services

import (
	"fmt"
	"time"
	"math/rand"
	"net"
)

type GameTimer struct {
	startTime time.Time
	maxDuration time.Duration
}

func NewGameTimer(maxMinutes int) *GameTimer {
	return &GameTimer{
		startTime: time.Now(),
		maxDuration: time.Duration(maxMinutes) * time.Minute,
	}
}

func (t *GameTimer) GetElapsedTime() string {
	elapsed := time.Since(t.startTime)
	return formatDuration(elapsed)
}

func (t *GameTimer) GetRemainingTime() string {
	remaining := t.maxDuration - time.Since(t.startTime)
	if remaining < 0 {
		remaining = 0
	}
	return formatDuration(remaining)
}

func (t *GameTimer) IsTimeUp() bool {
	return time.Since(t.startTime) >= t.maxDuration
}

func formatDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

func GenerateUniqueIPs(n int) []net.IP {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	ips := make([]net.IP, 0, n)
	seen := make(map[string]bool)

	for len(ips) < n {
		ip := make(net.IP, 4)
		for j := range ip {
			ip[j] = byte(rand.Intn(256))
		}
		ipStr := ip.String()
		if !seen[ipStr] {
			seen[ipStr] = true
			ips = append(ips, ip)
		}
	}
	return ips
}

// RemoveIP удаляет указанный IP из списка
func RemoveIP(ips []net.IP, target net.IP) []net.IP {
	for i, ip := range ips {
		if ip.Equal(target) {
			return append(ips[:i], ips[i+1:]...)
		}
	}
	return ips
}