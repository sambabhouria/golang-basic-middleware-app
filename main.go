// forms.go
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type ContactDetails struct {
	Name        string
	PhoneNumber string
	Email       string
	Options     string
	Message     string
}

func main() {
	tmpl := template.Must(template.ParseFiles("static/forms.html"))
	fmt.Println("hello world")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Name:        r.FormValue("name"),
			PhoneNumber: r.FormValue("phonenumber"),
			Email:       r.FormValue("email"),
			Options:     r.FormValue("options"),
			Message:     r.FormValue("message"),
		}
		fmt.Println("Forms Values", details)
		// do something with details
		_ = details

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.HandleFunc("/profile", logging(profile))
	http.HandleFunc("/fieldset", logging(bar))

	http.ListenAndServe(":8080", nil)
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func profile(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/profile.html"))
	tmpl.Execute(w, nil)
}

func bar(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/fieldset.html"))
	tmpl.Execute(w, nil)
}
