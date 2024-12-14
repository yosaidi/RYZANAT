package controllers

import (
	"html/template"
	"net/http"
)

var Tmpl *template.Template

type PageData struct {
	SuccessMessage string
}

func MainPage(w http.ResponseWriter, r *http.Request) {
if err := Tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}


func ErrorPage(w http.ResponseWriter, r *http.Request){
	if err := Tmpl.ExecuteTemplate(w, "error.html", nil); err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}

func LoginPage(w http.ResponseWriter, r *http.Request){
	if err := Tmpl.ExecuteTemplate(w, "login.html", nil); err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}

func Personalize(w http.ResponseWriter, r *http.Request){
	data := PageData{}
	if r.Method == http.MethodGet {
		name := r.URL.Query().Get("name")
		if name != "" {
			data.SuccessMessage = "Thank you, " + name + "! Your request has been submitted successfully."
		}
	}
	Tmpl.ExecuteTemplate(w, "personalize.html", data)
}