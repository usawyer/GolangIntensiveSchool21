package domain

import (
	"github.com/pkg/errors"
	"html/template"
	"net/http"
	"search/internal/db"
	"search/internal/types"
	"strconv"
)

type Server struct {
	store *db.ElasticStore
}

type Data struct {
	Places       []types.Place
	Total        int
	IsPrevious   bool
	IsNext       bool
	PreviousPage int
	NextPage     int
	LastPage     int
}

func NewServer(store *db.ElasticStore) *Server {
	return &Server{store: store}
}

func (server *Server) GetPlacesHandler(w http.ResponseWriter, r *http.Request) {
	data, err := server.PrepareData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("web/template/template.html"))
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}

func (server *Server) PrepareData(r *http.Request) (Data, error) {
	pageParam := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		return Data{}, errors.New("Invalid 'page' value " + pageParam)
	}

	limit := 10
	offset := (page - 1) * limit

	places, total, err := server.store.GetPlaces(limit, offset)
	if err != nil {
		return Data{}, errors.New("Error fetching places: " + err.Error())
	}

	answer := Data{
		Places: places,
		Total:  total,
	}

	if offset > 0 {
		answer.IsPrevious = true
		answer.PreviousPage = page - 1
	}

	if offset+limit < total {
		answer.IsNext = true
		answer.NextPage = page + 1
	}

	answer.LastPage = (total + limit - 1) / limit

	if page > answer.LastPage {
		return Data{}, errors.New("Invalid 'page' value " + pageParam)
	}

	return answer, nil
}
