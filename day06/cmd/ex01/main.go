package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
)

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/template/test.gohtml"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}

func adminPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ADMIN PAGE"))
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
