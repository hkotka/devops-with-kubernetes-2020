package main

import (
	"fmt"
	"net/http"
	"sync"
)

const (
	httpServePort = "8080"
)

type pongCount struct {
	count int
	mux   sync.Mutex
}

var pong pongCount

func main() {
	fmt.Println("Server started in port", httpServePort)
	http.HandleFunc("/pingpong", pingpong)
	http.HandleFunc("/pongcount", getPongCount)

	if err := http.ListenAndServe(":"+httpServePort, nil); err != nil {
		fmt.Println(err)
	}
}

func pingpong(w http.ResponseWriter, _ *http.Request) {
	pong.mux.Lock()
	defer pong.mux.Unlock()
	pong.count++
	resp := fmt.Sprintf("Ping / Pongs: %d", pong.count)
	fmt.Println(resp)

	if _, err := fmt.Fprintf(w, resp); err != nil {
		fmt.Println(err)
	}
}

func getPongCount(w http.ResponseWriter, _ *http.Request) {
	resp := fmt.Sprintf("%d", pong.count)
	if _, err := fmt.Fprintf(w, resp); err != nil {
		fmt.Println(err)
	}
}
