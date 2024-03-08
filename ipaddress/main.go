package main

import (
	"fmt"
	"regexp"
	"net"
)

func main() {
	text := "このテキストにはIPv4アドレスが含まれています。例えば、192.168.0.1や10.0.0.1などです。"

	// IPv4アドレスを抽出する正規表現パターン
	pattern := `(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`

	// 正規表現パターンにマッチするIPv4アドレスを取得
	regex := regexp.MustCompile(pattern)
	ipAddresses := regex.FindAllString(text, -1)

	// パースしたIPv4アドレスを表示
	for _, ipAddress := range ipAddresses {
		fmt.Println(ipAddress)
	}

	_, ipnet, err := net.ParseCIDR("10.0.0.0/24")
	if err != nil {
		fmt.Println("Do Not Parse")
	}
	ip1 := net.ParseIP("10.0.0.1")
	ip2 := net.ParseIP("10.1.0.1")
	fmt.Println(ipnet.Contains(ip1)) // true
	fmt.Println(ipnet.Contains(ip2)) // false

}

// Output
//❯ go run ./main.go
// 192.168.0.1
// 10.0.0.1
// true
// false
//