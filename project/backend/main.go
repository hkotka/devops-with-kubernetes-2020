package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"time"
)

type Todo struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var todos []Todo

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:30000", "http://localhost:5000", "https://localhost:30443/"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/todos", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"todos": todos,
		})
	})
	r.POST("/todos", func(c *gin.Context) {
		if c.Request.Body != nil {
			data, _ := ioutil.ReadAll(c.Request.Body)
			if err := addTodo(data); err != nil {
				log.Println(err)
			}
			c.JSON(200, gin.H{
				"message": "OK",
			})
		} else {
			c.JSON(500, gin.H{
				"message": "Empty HTTP request body",
			})
		}
	})
	if err := r.Run(); err != nil {
		fmt.Println(err)
	}
}

func addTodo(data []byte) error {
	var newTodo Todo
	if err := json.Unmarshal(data, &newTodo); err != nil {
		return err
	}
	fmt.Printf("Added new ToDo: %+v\n", newTodo)
	todos = append(todos, newTodo)
	debugListTodos()

	return nil
}

func debugListTodos() {
	fmt.Printf("All ToDo's:\n%+v\n", todos)
}
