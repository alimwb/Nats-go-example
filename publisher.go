package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	for i := 0; i < 5; i++ {
		message := fmt.Sprintf("message test %d", i+1)
		err = nc.Publish("test.subject", []byte(message))
		if err != nil {
			fmt.Println("err: ", err)
		} else {
			fmt.Println("send success ", message)
		}
		time.Sleep(2 * time.Second)
	}
}
