package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

const StaticURL = "/static/"

const StaticRoot = "static/"

type Context struct {
	Title  string
	Static string
}

func Home(w http.ResponseWriter, req *http.Request) {
	context := Context{Title: "Welcome!"}
	render(w, "index", context)
}

func Sources(w http.ResponseWriter, req *http.Request) {
	context := Context{Title: "Source Datastores"}
	render(w, "sources", context)
}

func Destinations(w http.ResponseWriter, req *http.Request) {
	context := Context{Title: "Destination Datastores"}
	render(w, "destinations", context)
}

func render(w http.ResponseWriter, tmpl string, context Context) {
	context.Static = StaticURL
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

func StaticHandler(w http.ResponseWriter, req *http.Request) {
	staticFile := req.URL.Path[len(StaticURL):]
	if len(staticFile) != 0 {
		f, err := http.Dir(StaticRoot).Open(staticFile)
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, req, staticFile, time.Now(), content)
			return
		}
	}
	http.NotFound(w, req)
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/sources/", Sources)
	http.HandleFunc("/destinations/", Destinations)
	http.HandleFunc(StaticURL, StaticHandler)
	err := http.ListenAndServe(":0", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
