package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/game/"):]

	fmt.Println("method:", r.Method, ", URL:", r.URL, ", Path:", path, ", RawPath:", r.URL.RawPath)

	p := &Page{}

	tmpl := path
	if len(path) == 0 {
		tmpl = "index.html"
	}

	fmt.Println("template:", tmpl)

	renderTemplate(w, tmpl, p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		fmt.Println("Templates Parse Files Error!", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		fmt.Println("Templates Execute Files Error!", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("SquirrelBattle1.0"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
