package main

import (
	"github.com/nats-io/stan.go"
	"strconv"
	"time"
)

func main() {
	sc, err := stan.Connect("prod", "clientID")
	if err != nil {
		panic(err)
	}
	defer sc.Close()

	for i := 0; i <= 100; i++ {
		sc.Publish("testing", []byte("Hello from iteration "+strconv.Itoa(i)))
		time.Sleep(2 * time.Second)
	}

}
