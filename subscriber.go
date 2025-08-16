package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	_, err = nc.Subscribe("test.subject", func(msg *nats.Msg) {
		fmt.Printf("recieved message: %s\n", string(msg.Data))
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("list")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
