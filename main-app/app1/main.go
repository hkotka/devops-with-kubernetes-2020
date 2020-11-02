package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	httpServePort = "8080"
	timestampFile = "/common-data/time.txt"
	pongFile      = "/common-data/pong.txt"
)

// Serves timestamp from a file and adds hash
func main() {
	fmt.Println("Server started in port", httpServePort)
	http.HandleFunc("/", defaultHandler)
	if err := http.ListenAndServe(":"+httpServePort, nil); err != nil {
		fmt.Println(err)
	}
}

func defaultHandler(w http.ResponseWriter, _ *http.Request) {
	timestamp, err := readFromFile(timestampFile)
	if err != nil {
		fmt.Println(err)
	}
	pongCount, err := readFromFile(pongFile)
	if err != nil {
		fmt.Println(err)
	}
	resp := timestamp + rndString() + "\n" + pongCount
	if _, err := fmt.Fprintf(w, resp); err != nil {
		fmt.Println(err)
	}
}

func rndString() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ0123456789-")
	strLenght := 48
	var s strings.Builder

	for i := 0; i < strLenght; i++ {
		s.WriteRune(chars[rand.Intn(len(chars))])
	}

	return s.String()
}

func readFromFile(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
