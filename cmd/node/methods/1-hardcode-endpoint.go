package methods

import (
	"log"

	"github.com/pebbe/zmq4"
)

func Hardcoded() {
	context, _ := zmq4.NewContext()
	socket, _ := context.NewSocket(zmq4.REQ)
	defer socket.Close()

	endpoint := "localhost:5555"

	println("Connecting to " + endpoint)
	err := socket.Connect("tcp://" + endpoint)
	if err != nil {
		log.Fatal(err)
	}

	socket.Send("Hi server :)", 0)
	socket.Recv(0)
}
