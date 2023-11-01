package subcriber

import (
	"github.com/pebbe/zmq4"
	"fmt"
)

func main() {
	subscriber, _ := zmq4.NewSocket(zmq4.SUB)
	defer subscriber.Close()
	subscriber.Connect("tcp://127.0.0.1:5555")
	subscriber.SetSubscribe("") // Subscribe to all messages

	for {
		msg, _ := subscriber.Recv(0)
		fmt.Printf("Menerima: %s\n", msg)
	}
}
