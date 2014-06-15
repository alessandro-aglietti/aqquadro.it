package hello

import (
//    "fmt"
    "net/http"
	"html/template"
	"time"
	"math/rand"
	"strconv"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    mux := http.NewServeMux()
	mux.HandleFunc("/api/", apiHandler)
	/*
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path != "/" {
				http.NotFound(w, req)
				return
			}
		fmt.Fprintf(w, "Welcome to the home page!")
	})
	*/
	
	mux.ServeHTTP(w, r)
}

// Prepare some data to insert into the template.
type Model struct {
	Titolo1, Titolo3 string
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	helloTemplate, err := template.ParseFiles("hello.html")
	
	sleepThis := rand.Intn(3000);
	time.Sleep(time.Duration(sleepThis) * time.Millisecond)
	
	err = helloTemplate.Execute(w, Model{"TITOLO 1", "Sleep for " + strconv.Itoa(sleepThis) + "ms"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}