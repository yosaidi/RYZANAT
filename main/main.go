package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"ryzanat/controllers"
	"ryzanat/router"
)

func main() {
	var err error
	style := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))

	http.Handle("/static/", style)
	
	controllers.Tmpl, err = template.ParseGlob("views/*.html")

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/",router.Router)
	fmt.Println("Server starts at  ---> http://localhost:8080/")
	
	http.ListenAndServe(":8080",nil)
}
