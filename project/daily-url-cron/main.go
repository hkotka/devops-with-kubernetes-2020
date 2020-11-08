package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

const url = "https://en.wikipedia.org/wiki/Special:Random"

type Todo struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
	gorm.Model
}

func main() {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable TimeZone=Europe/Helsinki",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_URL"), os.Getenv("POSTGRES_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connetion to Postgres successful")

	if todo, err := getWikiRndArticle(url); err != nil {
		log.Fatal(err)
	} else {
		err := db.Select("Name", "Done").Create(&todo).Error
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getWikiRndArticle(url string) (Todo, error) {
	var todo Todo

	resp, err := http.Get(url)
	if err != nil {
		return todo, err
	}
	defer resp.Body.Close()

	wikiUrl := fmt.Sprintf("%s://%s%s", resp.Request.URL.Scheme, resp.Request.URL.Host, resp.Request.URL.Path)

	if wikiUrl != "" {
		todo.Name = wikiUrl
		todo.Done = false
	} else {
		return todo, errors.New("wiki random request url empty")
	}

	return todo, nil
}
