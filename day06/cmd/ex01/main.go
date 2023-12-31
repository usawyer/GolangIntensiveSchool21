package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
)

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	redirect(w, r)

	tmpl := template.Must(template.ParseFiles("web/template/main.gohtml", "web/template/header.gohtml"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}

func adminPageHandler(w http.ResponseWriter, r *http.Request) {
	redirect(w, r)

	tmpl := template.Must(template.ParseFiles("web/template/admin.gohtml", "web/template/header.gohtml"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		action := r.FormValue("action")
		switch action {
		case "login":
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		case "back":
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		case "submit":
			login := r.FormValue("login")
			password := r.FormValue("password")

			if login != "" && password != "" {
				if login == "user" && password == "user" {
					log.Println("SUCCESS")
				} else {
					log.Println("WRONG")
				}
			}
		}

	}

}

func newRouter() *mux.Router {
	r := mux.NewRouter()

	staticFileDirectory := http.Dir("assets")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler)

	r.HandleFunc("/", mainPageHandler)
	r.HandleFunc("/admin", adminPageHandler)
	return r
}

func main() {
	r := newRouter()
	err := http.ListenAndServe("localhost:8888", r)
	if err != nil {
		log.Fatal(err)
	}
}
