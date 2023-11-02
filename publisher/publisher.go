package main

import (
	"fmt"
	"time"

	"github.com/pebbe/zmq4"
)

func main() {
	publisher, _ := zmq4.NewSocket(zmq4.PUB)
	defer publisher.Close()
	publisher.Bind("tcp://*:5555")

	messageQueue := make(chan string, 100) // Message queue untuk menyimpan pesan
	subscriberConnected := false

	// Fungsi untuk mengirim pesan dari message queue ke publisher
	sendMessagesFromQueue := func() {
		for {
			select {
			case message := <-messageQueue:
				publisher.Send(message, 0)
				fmt.Println("Published from queue:", message)
			}
		}
	}

	// Mulai goroutine untuk mengirim pesan dari message queue
	go sendMessagesFromQueue()

	for i := 0; ; i++ {
		message := fmt.Sprintf("Data %d", i)

		// Jika tidak ada subscriber yang terhubung, simpan pesan dalam message queue
		if !subscriberConnected {
			messageQueue <- message
		} else {
			if len(messageQueue) > 0 {
				// Kirim pesan dari message queue jika ada pesan yang tertunda
				sendMessagesFromQueue()
			}

			// Kirim pesan baru
			publisher.Send(message, 0)
			fmt.Println("Published:", message)
		}

		time.Sleep(time.Second)
	}
}
