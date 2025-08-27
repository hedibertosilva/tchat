package server

import (
	"log"
	"net"

	connection "github.com/hedibertosilva/tchat/internal/connection"
	consts "github.com/hedibertosilva/tchat/internal/consts"
)

var (
	Address     string
	NetListener net.Listener
)

func ConnListener() {
	NetListener, err := net.Listen(consts.Protocol, Address)
	if err != nil {
		log.Fatal(err)
	}

	connection.NetConn, err = NetListener.Accept()
	if err != nil {
		log.Fatal(err)
	}
}
