package main

import (
	"flag"
	"log"

	client "github.com/hedibertosilva/tchat/internal/client"
	connection "github.com/hedibertosilva/tchat/internal/connection"
	server "github.com/hedibertosilva/tchat/internal/server"
)

func handleArgs() {
	flag.StringVar(&server.Address, "listen", "", "Create Server.")          // Expected: host:port
	flag.StringVar(&client.Address, "connect", "", "Connect to the Server.") // Expected: host:port
	flag.Parse()
}

func main() {
	handleArgs()

	if server.Address != "" {
		log.Print("Listerning on ", server.Address)
		server.ConnListener()
	} else if client.Address != "" {
		log.Print("Connected to ", client.Address)
		client.ConnDialer()
	} else {
		log.Fatal("Please specify either --listen or --connect")
		return
	}

	defer connection.NetConn.Close()

	go connection.Sender()
	connection.Reader()
}
