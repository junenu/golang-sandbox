package main

import (
	"fmt"
	"regexp"
	"net"
	"os"
)

func main() {

	// IPv4アドレスを抽出する正規表現パターン
	pattern := `\b(?:\d{1,3}\.){3}\d{1,3}(?:\/\d{1,2})?\b`

	// sample/test.confの内容を読み込む
	file, err := os.Open("sample/test.conf")
	if err != nil {
		fmt.Println("Do Not Open")
	}
	defer file.Close()

	// ファイルの内容を取得
	buf := make([]byte, 1024)
	file.Read(buf)
	file.Close()
	fileStr := string(buf)
	//fmt.Println(fileStr)

	regex := regexp.MustCompile(pattern)
	ipAddresses := regex.FindAllString(fileStr, -1)

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