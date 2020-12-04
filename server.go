package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "loackhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string

var (
	entering = make(chan clinet)
	leaving  = make(chan clinet)
	message  = make(chan clinet)
)

func broadcaster() {
	client := make(map[clinet]bool)
	for {
		select {
		case msg := <-message:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			client[cli] = true
		case cli := <-leaving:
			delete(clinets, cli)
			close(cli)
		}
	}
}
