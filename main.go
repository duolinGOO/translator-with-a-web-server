package main

import (
	"fmt"
	"net"
)

var dict = map[string]string{
	"зеленый":    "къацу",
	"красный":    "яру",
	"синий":      "вили",
	"желтый":     "хъипи",
	"белый":      "лацу",
	"бесцветный": "рангсуз",
	"золотистый": "къизилдин",
	"коричневый": "шуьтруь",
	"оранжевый":  "туракь",
	"русый":      "расу",
	"серый":      "рехи",
	"смуглый":    "къумрав",
	"темный":     "мичIи",
	"черный":     "чIулав",
	"розовый":    "гити",
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server is listening")
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handleconnec(conn)
	}

}
func handleconnec(conn net.Conn) {
	for {
		defer conn.Close()
		buf := make([]byte, (1024 * 4))
		n, err := conn.Read(buf)
		if n == 0 || err != nil {
			fmt.Println("error:", err)
			break
		}
		if n == 0 || err != nil {
			fmt.Println("error: ", err)
			break
		}
		takestr := string(buf[0:n])
		answ, ok := dict[takestr]
		if ok == false {
			fmt.Println("not found")
		}
		conn.Write([]byte(answ))
	}
}
