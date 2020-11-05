package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

const (
	httpServePort = "8080"
)

type pongCount struct {
	Count int `json:"count"`
}

var pong pongCount
var mutex sync.Mutex

func main() {
	fmt.Println("Server started in port", httpServePort)
	http.HandleFunc("/pingpong", pingpong)
	http.HandleFunc("/pongcount", getPongCount)

	if err := http.ListenAndServe(":"+httpServePort, nil); err != nil {
		fmt.Println(err)
	}
}

func pingpong(w http.ResponseWriter, _ *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	pong.Count++
	resp := fmt.Sprintf("Ping / Pongs: %d", pong.Count)
	fmt.Println(resp)

	if _, err := fmt.Fprintf(w, resp); err != nil {
		fmt.Println(err)
	}
}

func getPongCount(w http.ResponseWriter, _ *http.Request) {
	resp, err := json.Marshal(pong)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if _, err := w.Write(resp); err != nil {
		fmt.Println(err)
	}
}
