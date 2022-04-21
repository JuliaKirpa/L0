package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"io/ioutil"
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
		sc.Publish("static", byteValue)
		fmt.Println("Push ", i)
		time.Sleep(5 * time.Second)
	}
}
