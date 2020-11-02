package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

const (
	httpServePort = "8080"
	fileLocation  = "/common-data/pong.txt"
)

type pongCount struct {
	count int
	mux   sync.Mutex
}

var pong pongCount

func main() {
	fmt.Println("Server started in port", httpServePort)
	http.HandleFunc("/", pingpong)
	http.ListenAndServe(":"+httpServePort, nil)
}

func pingpong(w http.ResponseWriter, _ *http.Request) {
	pong.mux.Lock()
	defer pong.mux.Unlock()
	pong.count++
	resp := fmt.Sprintf("Ping / Pongs: %d", pong.count)
	fmt.Println(resp)
	go func() {
		err := writeToFile(resp)
		if err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Fprintf(w, resp)
}

func writeToFile(pongCounter string) error {
	f, err := os.Create(fileLocation)
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(pongCounter)
	return nil
}
