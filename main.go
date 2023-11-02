package main

import (
	"fmt"
	"log"
	"time"

	"github.com/pebbe/zmq4"
)

func publisher() {
	publisher, _ := zmq4.NewSocket(zmq4.PUB)
	defer publisher.Close()
	publisher.Bind("tcp://*:5555")

	for {
		// Menerima pesan dari setiap subscriber
		_, err := publisher.Recv(0)
		if err != nil {
			log.Fatal(err)
		}

		// Mengirim pesan ke semua subscriber
		msg := "Hello, Subscribers!"
		publisher.Send(msg, zmq4.DONTWAIT)
		fmt.Printf("Sent: %s\n", msg)

		time.Sleep(1 * time.Second)
	}
}

func subscriber() {
	subscriber, _ := zmq4.NewSocket(zmq4.SUB)
	defer subscriber.Close()
	subscriber.Connect("tcp://localhost:5555")
	subscriber.SetSubscribe("") // Subscribe to all messages

	for {
		// Mengirim pesan kosong ke publisher agar terhitung sebagai subscriber yang terhubung
		subscriber.Send("", zmq4.SNDMORE)
		// Menerima pesan dari publisher
		msg, err := subscriber.Recv(0)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Received: %s\n", msg)
	}
}

func main() {
	go publisher()
	go subscriber()

	// Run indefinitely
	select {}
}
