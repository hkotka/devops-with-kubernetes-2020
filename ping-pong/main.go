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
	db    *gorm.DB
	Count int `json:"count" gorm:"default:0"`
}

func (p *PongCount) NewPong() {
	var err error
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable TimeZone=Europe/Helsinki",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_URL"), os.Getenv("POSTGRES_PORT"))

	p.Count = 0
	p.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err := p.db.AutoMigrate(&p); err != nil {
		log.Fatal(err)
	}
	err = p.db.First(&pong).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		pong.db.Save(&p)
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println(err)
	}
}

var pong PongCount

func main() {
	pong.NewPong()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:30000", "https://localhost:30443"},
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

func ginHandlerPongIncrement(c *gin.Context) {
	pong.db.First(&pong)
	pong.Count++
	if err := pong.db.Save(&pong).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	} else {
		c.String(200, "%d", pong.Count)
	}
}

func ginHandlerGetPongs(c *gin.Context) {
	if err := pong.db.First(&pong).Error; err != nil {
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
	c.JSON(200, gin.H{
		"health": "ok",
	})
}
