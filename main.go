package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"game_ip_wars/pkg/services"
	"game_ip_wars/pkg/handlers"
)

func main() {
	ips := services.GenerateUniqueIPs(10)
	handler := &handlers.CommandHandler{IPs: ips}
	scanner := bufio.NewScanner(os.Stdin)
	remaining := len(ips)
	attempts := 0
	timer := services.NewGameTimer(10) // 10 минут лимит

	fmt.Println("Добро пожаловать в игру IP Wars!")
	fmt.Printf("У вас есть 10 минут чтобы угадать %d уникальных IP-адресов\n\n", remaining)
	fmt.Println("Доступные команды:")
	fmt.Println("1. scan <маска> - найти IP по маске (например: scan 192.168)")
	fmt.Println("2. ping <IP> - проверить наличие IP в списке")
	fmt.Println("3. trace <IP> - найти соседние IP в сети")
	fmt.Println("Цель: найти и удалить все IP из списка")
	fmt.Printf("Всего IP: %d\n\n", len(ips))

	HelpIPList(ips)

	for len(handler.IPs) > 0 {
		fmt.Print("\nВведите команду: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		parts := strings.SplitN(input, " ", 2)

		if len(parts) < 1 {
			fmt.Println("Ошибка: неверная команда")
			continue
		}

		command := strings.ToLower(parts[0])
		var arg string
		if len(parts) > 1 {
			arg = parts[1]
		}

		switch command {
		case "scan":
			if arg == "" {
				fmt.Println("Ошибка: укажите маску для scan")
				continue
			}
			matched := handler.Scan(arg)
			if len(matched) > 0 {
				fmt.Println("Найденные IP:")
				for _, ip := range matched {
					fmt.Println("-", ip)
				}
			} else {
				fmt.Println("IP по указанной маске не найдены")
			}

		case "ping":
			if arg == "" {
				fmt.Println("Ошибка: укажите IP для ping")
				continue
			}
			if handler.Ping(arg) {
				fmt.Println("IP найден в списке!")
				ip := net.ParseIP(arg)
				handler.RemoveIP(ip)
				fmt.Printf("IP %s удален! Осталось: %d\n", ip, len(handler.IPs))
			} else {
				fmt.Println("IP не найден в списке")
			}

		case "trace":
			if arg == "" {
				fmt.Println("Ошибка: укажите IP для trace")
				continue
			}
			neighbors := handler.Trace(arg)
			if len(neighbors) > 0 {
				fmt.Println("Найденные соседние сети:")
				for _, n := range neighbors {
					fmt.Println("-", n)
				}
			} else {
				fmt.Println("Соседние сети не найдены")
			}

		default:
			fmt.Println("Неизвестная команда. Доступные команды: scan, ping, trace")
		}
	}

	fmt.Printf("\nИгра завершена! Затраченное время: %s\n", timer.GetElapsedTime())
	fmt.Printf("Угадано адресов: %d/%d\n", len(services.GenerateUniqueIPs(10))-remaining, len(services.GenerateUniqueIPs(10)))
	fmt.Printf("Всего попыток: %d\n", attempts)

	if remaining == 0 {
		fmt.Println("Поздравляем! Вы угадали все IP-адреса!")
	} else {
		fmt.Println("Попробуйте ещё раз, чтобы угадать оставшиеся адреса!")
	}
}

func HelpIPList (ips []net.IP){
	for i, ip := range ips {
		fmt.Printf("%d. %s\n", i+1, ip)
	}
}
