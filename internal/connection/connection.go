package connection

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var (
	NetConn net.Conn
)

func Sender() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		msg := scanner.Text()
		_, err := fmt.Fprintln(NetConn, msg)
		if err != nil {
			fmt.Println("Send error:", err)
			return
		}
	}
}

func Reader() {
	scanner := bufio.NewScanner(NetConn)

	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Println("Peer:", msg)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Read error:", err)
	}
}
