package client

import (
	"log"
	"net"

	connection "github.com/hedibertosilva/tchat/internal/connection"
	consts "github.com/hedibertosilva/tchat/internal/consts"
)

var (
	Address string
)

func ConnDialer() {
	var err error
	connection.NetConn, err = net.Dial(consts.Protocol, Address)
	if err != nil {
		log.Fatal(err)
	}
}
