package net

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

type Todo struct {
	Order int
	Is    bool
	Text  string
}

//var mainTemplate, _ = template.ParseFiles("html/main.html")
//var nextTemplate, _ = template.ParseFiles("html/next.html")

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s.", r.URL.Path[1:])
	mainTemplate, _ := template.ParseFiles("html/main.html")
	mainTemplate.Execute(w, nil)
}

func nextPage(w http.ResponseWriter, r *http.Request) {
	nextTemplate, _ := template.ParseFiles("html/next.html")
	nextTemplate.Execute(w, nil)
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Print("POST\n")
		res, _ := ioutil.ReadAll(r.Body)
		fmt.Printf("%s\n", res)
	}
}

func Call_http() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/next", nextPage)
	http.HandleFunc("/signup", signupHandler)
	http.ListenAndServe(":8080", nil)
}
