package server

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/mehmetolgundev/kvstore/server/command"
	"github.com/mehmetolgundev/kvstore/server/protocol"
)

type KVStore interface {
	Run()
}
type kvstore struct {
}

func NewKVStore() KVStore {
	return &kvstore{}
}

func (kv *kvstore) Run() {
	listener, err := net.Listen("tcp", "127.0.0.1:4250")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("KVstore is running.")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		var msg []byte
		for {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			msg = append(msg, buf[:n]...)
			if n < len(buf) {
				break
			}
		}
		isValid, cmd := protocol.Check(string(msg))
		if !isValid {
			conn.Write([]byte("Command is invalid"))
			conn.Close()
		}
		command.Do(cmd, string(msg))
		conn.Close()
	}

}
