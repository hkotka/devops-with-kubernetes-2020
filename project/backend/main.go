package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"unicode/utf8"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	ginLog "github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	todoLeghtLimit  = 140
	picsumUrl       = "https://picsum.photos/200"
	picDestLocation = "images/picsum.jpg"
)

var todoList Todos

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

func (t *Todo) String() string {
	return fmt.Sprintf("Added new todo wit ID: %d and task: %s", t.ID, t.Name)
}

func main() {
	todoList.Init()

	go func() {
		if !checkFileExists(picDestLocation) {
			if err := getNewPicsum(picsumUrl); err != nil {
				log.Println(err)
			}
		}
	}()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	ginLog.Logger = ginLog.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stdout,
			NoColor: false,
		},
	)

	r := gin.New()
	r.Use(logger.SetLogger())

	// Custom logger
	subLog := zerolog.New(os.Stdout).With().
		Logger()

	r.Use(logger.SetLogger(logger.Config{
		Logger: &subLog,
		UTC:    true,
	}))

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", ginHandlerGKEHealthcheck)
	r.Use(static.Serve("/images", static.LocalFile("/images", false)))
	r.GET("/todos", ginHandlerGetTodos)
	r.POST("/todos", ginHandlerPostTodo)

	if err := r.Run(); err != nil {
		fmt.Println(err)
	}
}

func ginHandlerGKEHealthcheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"health": "ok",
	})
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
	}
	if utf8.RuneCountInString(newTodo.Name) > todoLeghtLimit {
		msg := fmt.Sprintf("Rejected NewTodo: NewTodo had %d runes: Limit is %d", utf8.RuneCountInString(newTodo.Name), todoLeghtLimit)
		log.Println(msg)
		c.JSON(500, gin.H{
			"message": msg,
		})
		return
	} else {
		if err := todoList.db.Select("Name", "Done").Create(&newTodo).Error; err != nil {
			log.Println(err)
		} else {
			c.JSON(200, gin.H{
				"message": "OK",
			})
			log.Println(&newTodo)
		}
	}
}

func getNewPicsum(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	f, err := os.Create(picDestLocation)
	if err != nil {
		return err
	}
	if _, err = f.Write(data); err != nil {
		fmt.Println(err)
	}
	log.Printf("Fetched new Picsum image to %s\n", picDestLocation)

	return nil
}

func checkFileExists(file string) bool {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
