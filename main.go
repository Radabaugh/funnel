package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

const staticURL = "/static/"

const staticRoot = "static/"

type context struct {
	Static string
}

func home(w http.ResponseWriter, req *http.Request) {
	context := context{}
	render(w, "index", context)
}

func sources(w http.ResponseWriter, req *http.Request) {
	context := context{}
	render(w, "sources", context)
}

func destinations(w http.ResponseWriter, req *http.Request) {
	context := context{}
	render(w, "destinations", context)
}

func render(w http.ResponseWriter, tmpl string, context context) {
	context.Static = staticURL
	tmplList := []string{"templates/base.html",
		fmt.Sprintf("templates/%s.html", tmpl)}
	t, err := template.ParseFiles(tmplList...)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, context)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func staticHandler(w http.ResponseWriter, req *http.Request) {
	staticFile := req.URL.Path[len(staticURL):]
	if len(staticFile) != 0 {
		f, err := http.Dir(staticRoot).Open(staticFile)
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, req, staticFile, time.Now(), content)
			return
		}
	}
	http.NotFound(w, req)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/sources/", sources)
	http.HandleFunc("/destinations/", destinations)
	http.HandleFunc(staticURL, staticHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
