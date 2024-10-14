package methods

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pebbe/zmq4"
)

func Configuration() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	context, _ := zmq4.NewContext()
	socket, _ := context.NewSocket(zmq4.REQ)
	defer socket.Close()

	endpoint := os.Getenv("endpoint")

	println("Connecting to " + endpoint)
	err = socket.Connect("tcp://" + endpoint)
	if err != nil {
		log.Fatal(err)
	}

	socket.Send("Hi server :)", 0)
	socket.Recv(0)
}
