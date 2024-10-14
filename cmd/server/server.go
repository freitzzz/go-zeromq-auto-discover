package main

import (
	"github.com/pebbe/zmq4"
)

func main() {
	context, _ := zmq4.NewContext()
	socket, _ := context.NewSocket(zmq4.REP)
	defer socket.Close()
	endpoint := "*:5555"
	socket.Bind("tcp://" + endpoint)

	// Wait for messages
	println("Started server on endpoint: " + endpoint)
	for {
		msg, _ := socket.Recv(0)
		println("Received ", string(msg))

		// send reply back to client
		socket.Send("Hi node!!!", 0)
	}
}
