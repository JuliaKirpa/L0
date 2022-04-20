package main

import (
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"time"
)

func main() {
	sc, err := stan.Connect("prod", "static")
	if err != nil {
		panic(err)
	}
	defer sc.Close()

	byteValue, err := ioutil.ReadFile("./Producer/model.json")
	if err != nil {
		panic(err)
	}

	for i := 0; i <= 100; i++ {
		sc.Publish("static", byteValue)
		time.Sleep(3 * time.Second)
	}
}
