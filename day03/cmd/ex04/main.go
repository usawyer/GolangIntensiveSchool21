package main

import (
	"log"
	"net/http"
	"search/internal/db"
	"search/internal/domain"
)

func main() {
	store, err := db.NewElasticStore()
	if err != nil {
		log.Fatal(err)
	}

	server := domain.NewServer(store)

	http.HandleFunc("/api/get_token", domain.GetTokenHandler)
	http.HandleFunc("/api/recommend", domain.VerifyJWT(server.GetClosestPlacesHandler))

	err = http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}
}
