package main

import (
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type PongCount struct {
	gorm.Model
	Count int `json:"count" gorm:"default:0"`
}

var db *gorm.DB

func main() {
	db = initDb()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", ginHandlerGKEHealthcheck)
	r.GET("/pongcount", ginHandlerGetPongs)
	r.GET("/pingpong", ginHandlerPongIncrement)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}

func initDb() *gorm.DB {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable TimeZone=Europe/Helsinki",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_URL"), os.Getenv("POSTGRES_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var pong PongCount
	if err := db.AutoMigrate(&pong); err != nil {
		log.Fatal(err)
	}
	err = db.First(&pong).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		db.Save(&pong)
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println(err)
	}

	return db
}

func ginHandlerPongIncrement(c *gin.Context) {
	if err := db.Exec("UPDATE pong_counts SET count = count + 1;").Error; err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	} else {
		var pong PongCount
		if err := db.First(&pong).Error; err != nil {
			log.Println("Error getting new pongcount")
		} else {
			c.String(200, "%d", pong.Count)
		}
	}
}

func ginHandlerGetPongs(c *gin.Context) {
	var pong PongCount

	if err := db.First(&pong).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	} else {
		c.JSON(200, gin.H{
			"pongs": pong.Count,
		})
	}
}

func ginHandlerGKEHealthcheck(c *gin.Context) {
	var pong PongCount

	if err := db.First(&pong).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Database connection error",
		})
	} else {
		c.JSON(200, gin.H{
			"health": "ok",
		})
	}
}
