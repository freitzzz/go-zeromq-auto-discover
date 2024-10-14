package main

import (
	"log"

	"github.com/pebbe/zmq4"
	"github.com/zeromq/gyre"
)

func main() {
	context, _ := zmq4.NewContext()
	socket, _ := context.NewSocket(zmq4.REP)
	defer socket.Close()
	endpoint := "*:5555"
	socket.Bind("tcp://" + endpoint)

	go func() {
		startBeaconServer()
	}()

	// Wait for messages
	println("Started server on endpoint: " + endpoint)
	for {
		msg, _ := socket.Recv(0)
		println("Received ", string(msg))

		// send reply back to client
		socket.Send("Hi node!!!", 0)
	}
}

func startBeaconServer() {
	b, err := gyre.New()
	if err != nil {
		log.Fatal(err)
	}

	defer b.Stop()

	err = b.Start()
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case e := <-b.Events():
			switch e.Type() {
			case gyre.EventEnter:
				log.Printf("[%s] peer %q entered\n", b.Name(), e.Name())
				b.Whisper(e.Sender(), []byte("Hello"))
			case gyre.EventExit:
				log.Printf("[%s] peer %q exited\n", b.Name(), e.Name())
			default:
				println("exiting")
				return
			}
		}
	}
}
