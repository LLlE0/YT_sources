package main

import (
	"github.com/russross/blackfriday"
	"html/template"
	"log"
	"net/http"
)

func main() {
	//startin' message
	log.Println("The server is running!")

	//func handlers
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	})
	http.ListenAndServe(":8080", nil)
}

// third-party library to create a markdown✨
func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}
