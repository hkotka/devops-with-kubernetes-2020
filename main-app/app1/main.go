package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	httpServePort = "8080"
	timestampFile = "/common-data/time.txt"
	pongUrl       = "http://pingpong-svc:2346/pongcount"
)

type pong struct {
	Count int `json:"count"`
}

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
	pongCount, err := getPongCount(pongUrl)
	if err != nil {
		fmt.Println(err)
	}
	resp := timestamp + rndString() + "\n" + strconv.Itoa(pongCount)
	if _, err := fmt.Fprintf(w, resp); err != nil {
		fmt.Println(err)
	}
}

func getPongCount(url string) (int, error) {
	var pongcount pong

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil
	}
	if err = json.Unmarshal(data, &pongcount); err != nil {
		fmt.Println(err)
	}

	return pongcount.Count, nil
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
