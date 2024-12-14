package router

import (
	"net/http"
	"ryzanat/controllers"
)

func Router(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		controllers.MainPage(w, r)
	case "/login":
		controllers.LoginPage(w,r)
	case "/personalize":
		controllers.Personalize(w, r)
	default:
		controllers.ErrorPage(w, r)
	}

}
