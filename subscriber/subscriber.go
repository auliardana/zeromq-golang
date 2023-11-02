package main

import (
	"fmt"
	"log"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()

	// Hubungkan ke publisher
	err := subscriber.Connect("tcp://localhost:5555")
	if err != nil {
		log.Fatalf("Gagal menghubungkan subscriber: %s\n", err)
		return
	}

	subscriber.SetSubscribe("") // Subscribe to all messages

	for {
		message, err := subscriber.Recv(0)
		if err != nil {
			log.Printf("Gagal menerima pesan: %s\n", err)
			time.Sleep(1 * time.Second)
			continue
	}

	fmt.Println("Menerima: ", message)
	}
}
