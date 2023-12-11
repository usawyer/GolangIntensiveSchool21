package main

import (
	connector_db "go_team00/internal/connector-db"
	logger_create "go_team00/internal/logger-create"
	"go_team00/models"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Server struct {
	store  *gorm.DB
	logger *zap.Logger
}

func (s *Server) sessionsHandler(w http.ResponseWriter, r *http.Request) {
	var data []models.Session
	s.store.Find(&data)

	answer := models.SessionsModel{
		Count:    int64(len(data)),
		Sessions: data,
	}

	tmpl := template.Must(template.ParseFiles("web/template/sessions.gohtml"))
	if err := tmpl.Execute(w, answer); err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) anomalyHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data []models.Anomaly
	s.store.Where("session_id = ?", id).Find(&data)

	var session models.Session
	s.store.Where("id = ?", id).Find(&session)

	answer := models.AnomalyModel{
		Count:     int64(len(data)),
		Session:   session,
		Anomalies: data,
	}

	tmpl := template.Must(template.ParseFiles("web/template/anomaly.gohtml"))
	if err := tmpl.Execute(w, answer); err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	logger := logger_create.InitLogger("debug")
	store := connector_db.New(logger)

	data := Server{
		store:  store,
		logger: logger,
	}

	http.HandleFunc("/", data.sessionsHandler)
	http.HandleFunc("/anomaly", data.anomalyHandler)

	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
