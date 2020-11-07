package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Todos struct {
	todos []Todo
	db    *gorm.DB
}

type Todo struct {
	gorm.Model
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func (t *Todos) Init() {
	var err error
	if os.Getenv("POSTGRES_DB") != "" {
		dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable TimeZone=Europe/Helsinki",
			os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"),
			os.Getenv("POSTGRES_URL"), os.Getenv("POSTGRES_PORT"))
		t.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		log.Println("Connetion to Postgres successful")
	} else {
		t.db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
		log.Println("Connetion to postgres SQLlite")
	}
	if err != nil {
		panic("failed to connect database")
	}
	if err := t.db.AutoMigrate(&t.todos); err != nil {
		log.Fatal(err)
	}
}

var todoList Todos

func main() {
	todoList.Init()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:30000", "http://localhost:5000", "https://localhost:30443"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/todos", ginHandlerGetTodos)
	r.POST("/todos", ginHandlerPostTodo)

	if err := r.Run(); err != nil {
		fmt.Println(err)
	}
}

func ginHandlerGetTodos(c *gin.Context) {
	if err := todoList.db.Find(&todoList.todos).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	} else {
		c.JSON(200, gin.H{
			"todos": todoList.todos,
		})
	}
}

func ginHandlerPostTodo(c *gin.Context) {
	var newTodo Todo
	if c.Request.Body != nil {
		data, _ := ioutil.ReadAll(c.Request.Body)
		if err := json.Unmarshal(data, &newTodo); err != nil {
			log.Println(err)
			c.JSON(500, gin.H{
				"message": "Error, invalid json package",
			})
			return
		}
	} else {
		c.JSON(500, gin.H{
			"message": "Empty HTTP request body",
		})
	}
	if newTodo.Name == "" || newTodo.Done {
		log.Printf("Invalid ToDo json message. %+v", newTodo)
		c.JSON(500, gin.H{
			"message": "Error, invalid json package",
		})
		return
	} else if len(newTodo.Name) > 140 {
		c.JSON(500, gin.H{
			"message": "Too many characters. ToDo's can be up to 140 characters",
		})
		return
	} else {
		if err := todoList.db.Select("Name", "Done").Create(&newTodo).Error; err != nil {
			log.Println(err)
		} else {
			c.JSON(200, gin.H{
				"message": "OK",
			})
		}
	}
}
