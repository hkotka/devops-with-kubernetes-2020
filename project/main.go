package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	picsumUrl       = "https://picsum.photos/500"
	picDestLocation = "img/picsum.jpg"
	picsumLocalPath = "project/img/picsum.jpg"
	httpServePort   = "8080"
)

type TodoData struct {
	PageTitle string
	ImageLoc  string
	Todos     []Todo
}

type Todo struct {
	Title string
}

func main() {
	if err := getNewPicsum(picsumUrl); err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/project", defaultHandler)
	http.Handle("/project/img/", http.StripPrefix("/project/img/", http.FileServer(http.Dir("img"))))
	fmt.Println("Starting http server in port ", httpServePort)
	if err := http.ListenAndServe(":"+httpServePort, nil); err != nil {
		fmt.Println(err)
	}
}

func defaultHandler(w http.ResponseWriter, req *http.Request) {
	todo := TodoData{
		PageTitle: "MOOC Project",
		ImageLoc:  picsumLocalPath,
		Todos: []Todo{
			{Title: "TODO 1"},
			{Title: "TODO 2"},
		},
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, todo)
	log.Printf("%s - %s", req.Method, req.RequestURI)
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
