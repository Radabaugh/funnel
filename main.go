package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const staticURL = "/static/"

const staticRoot = "static/"

type context struct {
	Title  string
	Static string
}

func home(w http.ResponseWriter, req *http.Request) {
	context := context{Title: "Welcome!"}
	render(w, "index", context)
}

func sources(w http.ResponseWriter, req *http.Request) {
	context := context{Title: "Source Datastores"}
	render(w, "sources", context)
}

func destinations(w http.ResponseWriter, req *http.Request) {
	context := context{Title: "Destination Datastores"}
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

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {
	addr, e := determineListenAddress()
	if e != nil {
		log.Fatal(e)
	}

	http.HandleFunc("/", home)
	http.HandleFunc("/sources/", sources)
	http.HandleFunc("/destinations/", destinations)
	http.HandleFunc(staticURL, staticHandler)
	log.Printf("Listening on %s...\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
