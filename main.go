package main

import (
	"fmt"
	"net"

	"github.com/melsonic/gRedis/core"
	"github.com/melsonic/gRedis/util"
)

func main() {
	ln, lnServeError := net.Listen("tcp", util.ServerAddress)
	if lnServeError != nil {
		panic("Server crashed\n")
	}
	fmt.Printf("Server is listening on %d\n", util.PORT)
	defer ln.Close()
	for {
		conn, connErr := ln.Accept()
		if connErr != nil {
			panic("Error Setting TCP Connection\n")
		}
		go core.HandleConnection(conn)
	}
}
