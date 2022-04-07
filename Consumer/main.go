package main

import (
	"fmt"
	stan "github.com/nats-io/stan.go"
	"sync"
)

func main() {
	sc, err := stan.Connect("prod", "sub-1")
	if err != nil {
		panic(err)
	}
	defer sc.Close()
	wg := sync.WaitGroup{}
	wg.Add(1)
	// Simple Async Subscriber
	sub, err := sc.Subscribe("testing", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		panic(err)
	}

	defer sub.Unsubscribe()
	wg.Wait()
}
