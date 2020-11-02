package main

import (
	"fmt"
	"os"
	"time"
)

const timestampFile = "/common-data/time.txt"

var interval = time.Second * 5

// Generates timestamp to a file
func main() {
	for {
		t := timestamp()
		if err := writeToFile(t); err != nil {
			fmt.Println(err)
		}
		time.Sleep(interval)
	}
}

func timestamp() string {
	t := time.Now().Local()
	return fmt.Sprintf("%s", t)
}

func writeToFile(t string) error {
	f, err := os.Create(timestampFile)
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(t)
	return nil
}
