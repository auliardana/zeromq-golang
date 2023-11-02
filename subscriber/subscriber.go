package main

import (
	"fmt"
	"github.com/pebbe/zmq4"
)

func main() {
	subscriber, _ := zmq4.NewSocket(zmq4.SUB)
	defer subscriber.Close()
	subscriber.Connect("tcp://localhost:5555") // Sesuaikan dengan alamat publisher

	// Anda dapat menentukan filter untuk pesan yang akan diterima. Untuk menerima semua pesan, gunakan filter kosong.
	subscriber.SetSubscribe("")

	for {
		message, _ := subscriber.Recv(0)
		fmt.Println("Received:", message)
	}
}
