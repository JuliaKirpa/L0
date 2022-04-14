package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./client")))
	http.HandleFunc("/as", asyncHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func asyncHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	fmt.Fprintf(w, "data: %v\n\n", asc())

}

func asc() int {
	for i := 0; i <= 10; i++ {
		return rand.Intn(100)
		time.Sleep(2 * time.Second)
	}
	return -1
}
