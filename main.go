package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
	"math/rand"
	"game_ip_wars/pkg/services"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Генерация 10 случайных IP-адресов
	ips := services.generateIPs(10)
	targetIP := ips[rand.Intn(len(ips))]

	fmt.Println("Список IP-адресов:")
	for i, ip := range ips {
		fmt.Printf("%d. %s\n", i+1, ip)
	}

	fmt.Print("\nВведите IP-адрес, который вы считаете правильным: ")
	var input string
	fmt.Scanln(&input)

	// Нормализация ввода
	input = strings.TrimSpace(input)
	inputIP := net.ParseIP(input)

	if inputIP == nil {
		fmt.Println("Ошибка: введен некорректный IP-адрес")
		os.Exit(1)
	}

	if inputIP.Equal(targetIP) {
		fmt.Println("Поздравляем! Вы угадали правильный IP!")
	} else {
		fmt.Printf("К сожалению, вы не угадали. Правильный IP был: %s\n", targetIP)
	}
}
