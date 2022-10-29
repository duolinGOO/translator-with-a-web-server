package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Print("error: ", err)
		return
	}
	defer conn.Close()
	for {
		var origin string
		fmt.Print("Введите слово: ")
		_, err := fmt.Scanln(&origin)
		if err != nil {
			fmt.Println("try something else")
			continue
		}
		if n, err := conn.Write([]byte(origin)); n == 0 || err != nil {
			fmt.Println("error: ", err)
			return
		}
		fmt.Print("Перевод: ")
		bufforansw := make([]byte, 1024)
		n, err := conn.Read(bufforansw)
		if err != nil {
			break
		}
		fmt.Print(string(bufforansw[0:n]))
		fmt.Println()
	}
}
