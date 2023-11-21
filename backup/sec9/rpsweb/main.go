// building a simple webapp
package main

import (
	"log"
	"net/http"
	"text/template"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func main() {
	http.HandleFunc("/", homePage)

	// start a go built-in production ready web server
	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// http := `<h1>Hello World!</h1>`
	// w.Header().Set("Content-Type", "text/html")
	// fmt.Fprint(w, http)

}
