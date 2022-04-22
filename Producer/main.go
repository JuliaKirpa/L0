package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	sc, err := stan.Connect("prod", "cl-1")
	if err != nil {
		panic(err)
	}
	defer sc.Close()

	byteValue, err := ioutil.ReadFile("./Producer/model.json")
	if err != nil {
		panic(err)
	}

	for i := 0; i <= 100; i++ {
		err := sc.Publish("static", byteValue)
		if err != nil {
			log.Fatalf("can't publish mess: %s", err)
		}
		fmt.Println("Push ", i)
		time.Sleep(5 * time.Second)
	}
}
