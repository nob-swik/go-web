package main

import (
	"html/template"
	"log"
	"net/http"
)

// get target template
func page(fname string) *template.Template {
	tmps, _ := template.ParseFiles("templates/"+fname+".html", "templates/head.html", "templates/foot.html")
	return tmps
}

// index handler
func index(w http.ResponseWriter, rq *http.Request) {
	item := struct {
		Template string
		Title    string
		Message  string
	}{
		Template: "index",
		Title:    "Index",
		Message:  "This is Top page",
	}
	er := page("index").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

// hello handler
func hello(w http.ResponseWriter, rq *http.Request) {
	data := []string{
		"One", "Two", "Three",
	}

	item := struct {
		Title string
		Data  []string
	}{
		Title: "Hello",
		Data:  data,
	}

	er := page("hello").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		index(w, rq)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, rq *http.Request) {
		hello(w, rq)
	})

	http.ListenAndServe("localhost:8002", nil)
}
