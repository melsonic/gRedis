package core

import (
	"fmt"
	"io"
	"net"

	"github.com/melsonic/gRedis/util"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	var buffer []byte = util.EMPTY_READ_BYTE
	for {
		util.SetBytesToZero(buffer)
		_, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("reading done from " + conn.RemoteAddr().String() + " !!!")
			} else {
				fmt.Println("read err : ", err.Error())
			}
			break
		}
		requestArray, _ := util.Deserialize(buffer, []any{})
		response := Resolver(requestArray)
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Write Response Error:", err)
			break
		}
	}
}
