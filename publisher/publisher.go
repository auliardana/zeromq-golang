package publisher

import (
	"github.com/pebbe/zmq4"
	"fmt"
	"time"
)

func main() {
	publisher, _ := zmq4.NewSocket(zmq4.PUB)
	defer publisher.Close()
	publisher.Bind("tcp://127.0.0.1:5555")

	messageQueue := make(chan string, 100) // Buat channel untuk menyimpan pesan

	go func() {
		for {
			select {
			case msg := <-messageQueue:
				publisher.Send(msg, 0)
				fmt.Printf("Mengirim: %s\n", msg)
			}
		}
	}()

	for i := 0; ; i++ {
		msg := fmt.Sprintf("Data ke-%d", i)

		// Coba kirim pesan ke subscriber
		sent := publisher.Send(msg, zmq4.DONTWAIT)
		if sent == nil {
			fmt.Printf("Mengirim: %s\n", msg)
		} else {
			// Jika tidak bisa mengirim (subscriber tidak terhubung), simpan pesan di antrean
			messageQueue <- msg
		}

		time.Sleep(time.Second) // Menunggu sejenak sebelum mengirim pesan berikutnya
	}
}
