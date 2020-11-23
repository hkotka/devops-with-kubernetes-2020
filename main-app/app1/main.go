package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
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
	Count int `json:"pongs"`
}

// Serves timestamp from a file and adds hash
func main() {
	fmt.Println("Server started in port", httpServePort)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/healthz", ginHandlerGKEHealthcheck)
	r.GET("/", ginDefaultHandler)
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}

func ginDefaultHandler(c *gin.Context) {
	msg := os.Getenv("MESSAGE")
	timestamp, err := readFromFile(timestampFile)
	if err != nil {
		fmt.Println(err)
	}
	pongCount, err := getPongCount(pongUrl)
	if err != nil {
		fmt.Println(err)
	}
	c.String(200, "%s", msg+"\n"+timestamp+rndString()+"\n"+strconv.Itoa(pongCount))
}

func ginHandlerGKEHealthcheck(c *gin.Context) {
	if _, err := getPongCount(pongUrl); err != nil {
		c.JSON(500, gin.H{
			"health": err,
		})
	} else {
		c.JSON(200, gin.H{
			"health": "ok",
		})
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
