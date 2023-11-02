package main

import (
	"fmt"
	"log"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	publisher, _ := zmq.NewSocket(zmq.PUB)
	defer publisher.Close()

	// Bind publisher ke alamat endpoint
	err := publisher.Bind("tcp://*:5555")
	if err != nil {
		log.Fatalf("Gagal mengikat publisher: %s\n", err)
	}

	// Simpan pesan dalam slice
	messageQueue := []string{}
	i := 1 // Inisialisasi nilai i

	// Aktifkan publisher.monitor
	monitorEndpoint := "inproc://monitor"
	err = publisher.Monitor(monitorEndpoint, zmq.EVENT_CONNECTED|zmq.EVENT_DISCONNECTED)
	if err != nil {
		log.Fatalf("Gagal mengaktifkan monitor: %s\n", err)
	}

	go func() {
		monitor, _ := zmq.NewSocket(zmq.PAIR)
		defer monitor.Close()
		err := monitor.Connect(monitorEndpoint)
		if err != nil {
			log.Fatalf("Gagal menghubungkan monitor: %s\n", err)
		}

		for {
			event, addr, _, _ := monitor.RecvEvent(0) // Mengabaikan nilai keempat
			if event == zmq.EVENT_CONNECTED {
				log.Printf("Subscriber terhubung ke: %s\n", addr)
				// Kirim pesan yang ada di messageQueue saat subscriber terhubung
				for len(messageQueue) > 0 {
					message := messageQueue[0]
					publisher.Send(message, 0)
					messageQueue = messageQueue[1:]
				}
			} else if event == zmq.EVENT_DISCONNECTED {
				log.Printf("Subscriber terputus dari: %s\n", addr)
			}
		}
	}()

	for {
		// Kirim pesan ke subscriber
		message := fmt.Sprintf("Pesan ke %d", i)
		fmt.Println("Mengirim: ", message)
		i++

		// Simpan pesan dalam slice, bahkan jika tidak ada subscriber terhubung
		messageQueue = append(messageQueue, message)

		// Cobalah mengirim pesan ke subscriber
		_, err := publisher.Send(message, zmq.DONTWAIT)
		if err != nil {
			log.Println("Subscriber tidak terhubung. Menunggu...")
		} else {
			// Hapus pesan yang sudah terkirim dari slice
			messageQueue = messageQueue[1:]
		}

		time.Sleep(1 * time.Second)
	}
}
